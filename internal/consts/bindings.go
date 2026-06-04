package consts

import (
	"gta/internal/config"
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

var (
	LanguageBinding         = binding.NewString()
	SelectedModelBinding    = binding.NewInt()
	TranslationFromBinding  = binding.NewString()
	TranslationToBinding    = binding.NewString()
	GPRCHostBinding         = binding.NewString()
	GPRCPortBinding         = binding.NewInt()
	ModelURLBinding         = binding.NewString()
	ModelTemperatureBinding = binding.NewFloat()
	ModelTopPBinding        = binding.NewFloat()
	ModelTopKBinding        = binding.NewInt()
	ModelMaxTokensBinding   = binding.NewInt()
)

func InitBindingVariables(c *config.Config) error {
	a := fyne.CurrentApp()

	from := a.Preferences().StringWithFallback(CurrentFromLangKey, c.Translation.Source)
	err := TranslationFromBinding.Set(from)
	if err != nil {
		slog.Error("Failed to set TranslationFromBinding", "error", err)
	}

	to := a.Preferences().StringWithFallback(CurrentToLangKey, c.Translation.Target)
	err = TranslationToBinding.Set(to)
	if err != nil {
		slog.Error("Failed to set TranslationToBinding", "error", err)
	}

	h := a.Preferences().StringWithFallback(CurrentPortKey, c.GRPC.Host)
	err = GPRCHostBinding.Set(h)
	if err != nil {
		slog.Error("Failed to set GPRCHostBinding", "error", err)
	}

	p := a.Preferences().IntWithFallback(CurrentPortKey, c.GRPC.Port)
	err = GPRCPortBinding.Set(p)
	if err != nil {
		slog.Error("Failed to set GPRCPortBinding", "error", err)
	}

	url := a.Preferences().StringWithFallback(CurrentModelKey, c.Model.URL)
	err = ModelURLBinding.Set(url)
	if err != nil {
		slog.Error("Failed to set ModelURLBinding", "error", err)
	}

	temp := a.Preferences().FloatWithFallback(CurrentTemperatureKey, c.Model.Temperature)
	err = ModelTemperatureBinding.Set(temp)
	if err != nil {
		slog.Error("Failed to set ModelTemperatureBinding", "error", err)
	}

	topP := a.Preferences().FloatWithFallback(CurrentTopPKey, c.Model.TopP)
	err = ModelTopPBinding.Set(topP)
	if err != nil {
		slog.Error("Failed to set ModelURLBinding", "error", err)
	}

	topK := a.Preferences().IntWithFallback(CurrentTopKKey, c.Model.TopK)
	err = ModelTopKBinding.Set(topK)
	if err != nil {
		slog.Error("Failed to set ModelTopKBinding", "error", err)
	}

	maxTokens := a.Preferences().IntWithFallback(CurrentMaxTokensKey, c.Model.MaxTokens)
	err = ModelMaxTokensBinding.Set(maxTokens)
	if err != nil {
		slog.Error("Failed to set ModelMaxTokensBinding", "error", err)
	}

	return err
}

func InitBindingVariables2(c *config.Config) (err error) {
	a := fyne.CurrentApp()

	a.Preferences().SetString(CurrentFromLangKey, c.Translation.Source)
	err = TranslationFromBinding.Set(c.Translation.Source)
	if err != nil {
		slog.Error("Failed to set TranslationFromBinding", "error", err)
	}

	a.Preferences().SetString(CurrentToLangKey, c.Translation.Target)
	err = TranslationToBinding.Set(c.Translation.Target)
	if err != nil {
		slog.Error("Failed to set TranslationToBinding", "error", err)
	}

	a.Preferences().SetString(CurrentModelKey, c.Model.URL)
	err = ModelURLBinding.Set(c.Model.URL)
	if err != nil {
		slog.Error("Failed to set ModelURLBinding", "error", err)
	}

	a.Preferences().SetFloat(CurrentTemperatureKey, c.Model.Temperature)
	err = ModelTemperatureBinding.Set(c.Model.Temperature)
	if err != nil {
		slog.Error("Failed to set ModelTemperatureBinding", "error", err)
	}

	a.Preferences().SetFloat(CurrentTopPKey, c.Model.TopP)
	err = ModelTopPBinding.Set(c.Model.TopP)
	if err != nil {
		slog.Error("Failed to set ModelURLBinding", "error", err)
	}

	a.Preferences().SetInt(CurrentTopKKey, c.Model.TopK)
	err = ModelTopKBinding.Set(c.Model.TopK)
	if err != nil {
		slog.Error("Failed to set ModelTopKBinding", "error", err)
	}

	a.Preferences().SetInt(CurrentMaxTokensKey, c.Model.MaxTokens)
	err = ModelMaxTokensBinding.Set(c.Model.MaxTokens)
	if err != nil {
		slog.Error("Failed to set ModelMaxTokensBinding", "error", err)
	}

	return err
}
