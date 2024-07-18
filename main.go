package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/revett/miniflux-sync/cmd"
	"github.com/revett/miniflux-sync/config"
	"github.com/urfave/cli/v2"
)

// TODO: Improve test coverage.
// TODO: Add CI.
// TODO: Add README.
// TODO: Add Dependabot.
// TODO: Define release process.
// TODO: Add logger.

//go:embed VERSION
var version string

func main() {
	cfg := config.New(version)

	app := &cli.App{
		Name:     "miniflux-sync",
		Usage:    "Manage and sync your Miniflux feeds with YAML.",
		Version:  cfg.Version,
		Flags:    cfg.Flags(),
		Commands: cmd.Commands(cfg),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
