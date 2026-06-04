package gui

import (
	"gta/internal/config"
	"gta/internal/consts"
	"gta/internal/model"
	"gta/internal/ui"
	"gta/resource"
	"log/slog"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/lang"
)

var (
	ma model.AppObject
)

func Run(conf *config.Config) (err error) {
	// Set the scale of the app to 1.0
	os.Setenv("FYNE_SCALE", "1.0")
	slog.Debug("Locale", "lang", lang.SystemLocale())

	ma.App = app.NewWithID("cn.osdir.gta")
	ma.Config = conf
	lang.AddTranslationsFS(resource.Translations, "translation")

	ma.Window = ma.App.NewWindow(lang.L("gta_translation_assistant"))
	ma.Window.Resize(fyne.NewSize(800, 600))

	err = consts.InitBindingVariables(conf)
	if err != nil {
		slog.Error("Failed to set binding variables", "error", err)
		// os.Exit(1)
	}

	ma.Window.SetContent(ui.Create(ma))
	ma.Window.ShowAndRun()

	return
}
