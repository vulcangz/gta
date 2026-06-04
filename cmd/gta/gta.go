package gta

import (
	"context"
	"fmt"
	pb "gta/api/comm/v1"
	"gta/internal/config"
	"log/slog"
	"net"
	"strings"
	"time"

	"github.com/ardanlabs/kronk/sdk/kronk"
	"github.com/ardanlabs/kronk/sdk/kronk/model"
	"github.com/ardanlabs/kronk/sdk/tools/libs"
	"github.com/ardanlabs/kronk/sdk/tools/models"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// server is used to implement EchoServer.
type server struct {
	pb.UnimplementedGtrServer
	client pb.GtrClient
	cc     *grpc.ClientConn

	krnChat *kronk.Kronk
	conf    *config.Config
}

func newGtrServer(conf *config.Config) *server {
	address := fmt.Sprintf("%v:%v", conf.GRPC.Host, conf.GRPC.Port)
	cc, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("newGtrServer", "did not connect: ", err)
	}
	return &server{
		client: pb.NewGtrClient(cc),
		cc:     cc,
		conf:   conf}
}

func GTA(conf *config.Config) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	libs, err := libs.New()
	if err != nil {
		slog.Error("install-system", "unable to create libs api:", err)
		return
	}

	_, err = libs.Download(ctx, kronk.FmtLogger)
	if err != nil {
		slog.Error("install-system", "unable to install llama.cpp:", err)
		return
	}

	mdls, err := models.New()
	if err != nil {
		slog.Error("model", "unable to create models api:", err)
		return
	}

	modelChatURL := conf.Model.URL
	infoChat, err := mdls.Download(context.Background(), kronk.FmtLogger, modelChatURL)
	if err != nil {
		slog.Error("model", "unable to install model:", err)
		return
	}

	// -------------------------------------------------------------------------
	slog.Info("loading model...")
	if err = kronk.Init(); err != nil {
		slog.Error("model", "unable to init kronk:", err)
		return
	}

	cfg := model.Config{
		ModelFiles: infoChat.ModelFiles,
	}

	krnChat, err := kronk.New(model.WithConfig(cfg))
	if err != nil {
		slog.Error("model", "unable to create chat model: ", err)
	}
	defer func() {
		if err := krnChat.Unload(context.Background()); err != nil {
			slog.Error("model", "failed to unload chat model: ", err)
		}
	}()

	slog.Info("system info\n")
	for k, v := range krnChat.SystemInfo() {
		slog.Info("- system:", k, v)
	}
	slog.Info("\n")

	slog.Info("model info\n")
	slog.Info("- model", slog.Int("contextWindow", krnChat.ModelConfig().ContextWindow()))
	slog.Info("- model", slog.Bool("embeddings", krnChat.ModelInfo().IsEmbedModel))
	slog.Info("- model", slog.Bool("isGPT", krnChat.ModelInfo().IsGPTModel))
	slog.Info("\n")

	gtrServer := newGtrServer(conf)
	defer gtrServer.Close()
	gtrServer.krnChat = krnChat

	grpcServer := grpc.NewServer()
	pb.RegisterGtrServer(grpcServer, gtrServer)

	address := fmt.Sprintf("%v:%v", conf.GRPC.Host, conf.GRPC.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to listen: %v", err))
	}

	if err := grpcServer.Serve(lis); err != nil {
		slog.Error(fmt.Sprintf("failed to serve: %v", err))
	}

	slog.Info("startup: gRPC server running on ", slog.String("address", address))

	return
}

func (s *server) Translate(req *pb.TranslateRequest, stream pb.Gtr_TranslateServer) error {
	traceID := uuid.NewString()

	slog.Info("traceID: ", traceID, " translating started\n")
	defer slog.Info("traceID: ", traceID, " translating started\n")

	// 20*time.Minute
	callCtx, cancelCall := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancelCall()

	callCtx = kronk.SetFmtLoggerTraceID(callCtx, traceID)

	d := model.D{
		"messages": model.DocumentArray(
			model.TextMessage(model.RoleUser, req.Text),
		),
		"temperature": s.conf.Model.Temperature,
		"top_p":       s.conf.Model.TopP,
		"top_k":       s.conf.Model.TopK,
		"max_tokens":  s.conf.Model.MaxTokens,
	}

	ch, err := s.krnChat.ChatStreaming(callCtx, d)
	if err != nil {
		return fmt.Errorf("chat streaming: %w", err)
	}

	// -------------------------------------------------------------------------

	var (
		reasoning bool
		content   strings.Builder
		lr        model.ChatResponse
	)

	for resp := range ch {
		lr = resp

		switch resp.Choices[0].FinishReason() {
		case model.FinishReasonError:
			return fmt.Errorf("error from model: %s", resp.Choices[0].Delta.Content)

		case model.FinishReasonStop:
			return nil

		default:
			if resp.Choices[0].Delta.Reasoning != "" {
				reasoning = true
				fmt.Printf("\u001b[91m%s\u001b[0m", resp.Choices[0].Delta.Reasoning)
				continue
			}

			if reasoning {
				reasoning = false
				fmt.Println()
				continue
			}

			content.WriteString(resp.Choices[0].Delta.Content)

			resp := &pb.TranslateResponse{Message: resp.Choices[0].Delta.Content}
			if err := stream.Send(resp); err != nil {
				return status.Errorf(codes.Internal, "error sending stream: %v", err)
			}

		}
	}

	contextTokens := lr.Usage.PromptTokens + lr.Usage.CompletionTokens
	contextWindow := s.krnChat.ModelConfig().ContextWindow()
	percentage := (float64(contextTokens) / float64(contextWindow)) * 100
	of := float32(contextWindow) / float32(1024)
	slog.Info(fmt.Sprintf("\n\n\u001b[90mInput: %d  Reasoning: %d  Completion: %d  Output: %d  Window: %d (%.0f%% of %.0fK) TPS: %.2f\u001b[0m\n",
		lr.Usage.PromptTokens, lr.Usage.ReasoningTokens, lr.Usage.CompletionTokens, lr.Usage.OutputTokens, contextTokens, percentage, of, lr.Usage.TokensPerSecond))

	return nil
}

func (s *server) Close() {
	s.cc.Close()
}
