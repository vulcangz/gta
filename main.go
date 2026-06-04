package main

import (
	"gta/cmd/gta"
	"gta/cmd/gui"
	"gta/internal/config"

	"log/slog"
	"os"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {
	var (
		err  error
		conf *config.Config
	)

	conf, err = config.Load()
	if err != nil {
		slog.Error("startup", "ERROR", err)
		os.Exit(1)
	}

	arg0 := conf.Args.Num(0)
	switch arg0 {
	case "gui":
		// Perform the serve startup and shutdown sequence.
		if err := gui.Run(conf); err != nil {
			slog.Error("startup", "ERROR", err)
			os.Exit(1)
		}
	case "serve":
		fallthrough
	default:
		// Perform the serve startup and shutdown sequence.
		if err := gta.GTA(conf); err != nil {
			slog.Error("startup", "ERROR", err)
			os.Exit(1)
		}
	}
}
