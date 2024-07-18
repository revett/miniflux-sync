package config

import "github.com/urfave/cli/v2"

// SyncFlags holds the flags for the sync command.
type SyncFlags struct {
	Path string
}

// Flags returns the flags for the sync command.
func (s *SyncFlags) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Usage:       "Path to file for imported data. (required)",
			EnvVars:     []string{"MINIFLUX_SYNC_PATH"},
			Destination: &s.Path,
			Aliases:     []string{"p"},
			Required:    true,
		},
	}
}
