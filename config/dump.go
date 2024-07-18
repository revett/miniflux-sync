package config

import "github.com/urfave/cli/v2"

// DumpFlags holds the flags for the dump command.
type DumpFlags struct {
	Path string
}

// Flags returns the flags for the dump command.
func (d *DumpFlags) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Usage:       "Path to file for exported data. (optional)",
			EnvVars:     []string{"MINIFLUX_SYNC_PATH"},
			Destination: &d.Path,
			Aliases:     []string{"p"},
		},
	}
}
