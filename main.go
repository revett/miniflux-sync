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
// TODO: Define release process.
// TODO: Check that input files exist.
// TODO: Export remote as YAML.
// TODO: Bug - referencing category that was just created when creating a feed.

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


INF connecting to miniflux instance
INF checking health of miniflux instance
INF reading data from yaml file
INF ./examples/feeds.yml
INF local feeds count=7
INF local categories count=3
INF fetching feeds
INF fetching categories
INF remote feeds count=6
INF remote categories count=3
INF actions to perform count=1
INF createfeed category_title=Blog feed_url=https://matt-rickard.com/rss
INF performing actions
INF creating feed category=Blog url=https://matt-rickard.com/rss
