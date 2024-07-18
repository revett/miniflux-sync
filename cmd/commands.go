package cmd

import (
	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/config"
	"github.com/urfave/cli/v2"
)

// Commands returns the commands for the CLI.
func Commands(cfg *config.Config) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "sync",
			Aliases: []string{"s"},
			Usage:   "Update Miniflux using a local YAML or OPML file.",
			Action: func(ctx *cli.Context) error {
				if err := sync(cfg); err != nil {
					return errors.Wrap(err, "running sync command")
				}

				return nil
			},
		},
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "Dump the current remote Miniflux state to your machine.",
			Action: func(ctx *cli.Context) error {
				if err := dump(cfg); err != nil {
					return errors.Wrap(err, "running dump command")
				}

				return nil
			},
		},
	}
}
