package model

import (
	"gta/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// AppObject holds the main app and channels
type AppObject struct {
	App            fyne.App
	Window         fyne.Window
	Config         *config.Config
	InputBox       *widget.Entry
	OutputBox      *widget.Entry
	sourceLanguage *string
	targetLanguage *string
}

func (a *AppObject) SetSourceLanguage(s string) {
	a.sourceLanguage = &s
}

func (a *AppObject) GetSourceLanguage() *string {
	return a.sourceLanguage
}

func (a *AppObject) SetTargetLanguage(s string) {
	a.targetLanguage = &s
}

func (a *AppObject) GetTargetLanguage() *string {
	return a.targetLanguage
}
