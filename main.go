package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/sync"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "miniflux-sync",
		Usage: "Manage and sync your Miniflux feeds with YAML. ",
		Action: func(ctx *cli.Context) error {
			if err := sync.Sync(); err != nil {
				return errors.Wrap(err, "syncing config")
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %+v\n", err)
		os.Exit(1)
	}
}
