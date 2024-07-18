package config

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

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
			Action: func(ctx *cli.Context, s string) error {
				ext := filepath.Ext(s)

				if ext != ".xml" {
					log.Printf(`invalid file extension: "%s"`, ext)
					log.Printf(`allowed extension: ".xml"`)
					return errors.New("invalid file extension")
				}

				return nil
			},
		},
	}
}
