package config

import (
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/ardanlabs/conf/v3"
)

var mu sync.Mutex

const (
	appEnvPrefix = "GTA"
)

type LogLevel string

const (
	Debug    LogLevel = "debug"
	Info     LogLevel = "info"
	Warning  LogLevel = "warning"
	Critical LogLevel = "critical"
)

type Config struct {
	conf.Version
	Args        conf.Args
	AppName     string `conf:"default:GTA"`
	LogLevel    LogLevel
	Model       Model
	Translation Translation
	GRPC        GRPC
}

type Translation struct {
	Languages  []string `conf:"default:English;中文;Français;Italiano;日本語;한국어;Deutsch;繁體中文"`
	InputDelay string   `conf:"default:300"`
	Source     string   `conf:"default:English"`
	Target     string   `conf:"default:中文"`
}

type Model struct {
	URL         string  `conf:"default:https://huggingface.co/unsloth/gemma-4-E2B-it-GGUF/blob/main/gemma-4-E2B-it-Q4_K_M.gguf"`
	Temperature float64 `conf:"default:0.7"`
	TopP        float64 `conf:"default:0.9"`
	TopK        int     `conf:"default:40"`
	MaxTokens   int     `conf:"default:2048"`
}

type GRPC struct {
	Host string `conf:"default:0.0.0.0"`
	Port int    `conf:"default:9000"`
}

// LoadConfig returns the configuration
func Load() (*Config, error) {
	var cfg Config
	if h, err := conf.Parse(appEnvPrefix, &cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(h)
		}
		return nil, err
	}

	c, err := conf.String(&cfg)
	if err != nil {
		return nil, err
	}

	slog.Debug("Config", slog.String("data", c))

	return &cfg, nil
}
