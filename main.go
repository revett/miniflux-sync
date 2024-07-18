package main

import (
	"context"
	_ "embed"
	"os"

	"github.com/revett/miniflux-sync/cmd"
	"github.com/revett/miniflux-sync/config"
	"github.com/revett/miniflux-sync/log"
	zerolog "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// TODO: Improve test coverage.

//go:embed VERSION
var version string

func main() {
	ctx := context.Background()

	// Create logger, and attach to context.
	zerolog.Logger = log.New()
	ctx = zerolog.With().Logger().WithContext(ctx)

	cfg := config.New(version)

	app := &cli.App{
		Name:     "miniflux-sync",
		Usage:    "Manage and sync your Miniflux feeds with YAML.",
		Version:  cfg.Version,
		Flags:    cfg.Flags(),
		Commands: cmd.Commands(ctx, cfg),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(ctx, err)
	}
}
