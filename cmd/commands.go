package cmd

import (
	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/api"
	"github.com/revett/miniflux-sync/config"
	"github.com/urfave/cli/v2"
)

// Commands returns the commands for the CLI.
func Commands(cfg *config.GlobalFlags) []*cli.Command {
	dumpFlags := &config.DumpFlags{}
	syncFlags := &config.SyncFlags{}

	return []*cli.Command{
		{
			Name:    "sync",
			Aliases: []string{"s"},
			Usage:   "Update Miniflux using a local YAML file.",
			Flags:   syncFlags.Flags(),
			Action: func(ctx *cli.Context) error {
				client, err := api.Client(cfg)
				if err != nil {
					return errors.Wrap(err, "creating miniflux client")
				}

				if err := sync(syncFlags, client); err != nil {
					return errors.Wrap(err, "running sync command")
				}

				return nil
			},
		},
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "Dump the current remote Miniflux state to your machine.",
			Flags:   dumpFlags.Flags(),
			Action: func(ctx *cli.Context) error {
				client, err := api.Client(cfg)
				if err != nil {
					return errors.Wrap(err, "creating miniflux client")
				}

				if err := dump(dumpFlags, client); err != nil {
					return errors.Wrap(err, "running dump command")
				}

				return nil
			},
		},
	}
}
