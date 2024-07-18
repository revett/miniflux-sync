package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/revett/miniflux-sync/cmd"
	"github.com/revett/miniflux-sync/config"
	"github.com/urfave/cli/v2"
)

//go:embed VERSION
var version string

func main() {
	cfg := config.New(version)

	app := &cli.App{
		Name:     "miniflux-sync",
		Usage:    "Manage and sync your Miniflux feeds with YAML or OPML.",
		Version:  cfg.Version,
		Flags:    cfg.Flags(),
		Commands: cmd.Commands(cfg),
	}

	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %+v\n", err)
		os.Exit(1)
	}
}
