package config

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

// SyncFlags holds the flags for the sync command.
type SyncFlags struct {
	DryRun bool
	Path   string
}

// Flags returns the flags for the sync command.
func (s *SyncFlags) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "dry-run",
			Usage:       "Perform a dry run without making any changes.",
			EnvVars:     []string{"MINIFLUX_SYNC_DRY_RUN"},
			Destination: &s.DryRun,
			Aliases:     []string{"d"},
			Value:       false,
		},
		&cli.StringFlag{
			Name:        "path",
			Usage:       "Path to file for imported data. (required)",
			EnvVars:     []string{"MINIFLUX_SYNC_PATH"},
			Destination: &s.Path,
			Aliases:     []string{"p"},
			Required:    true,
			Action: func(ctx *cli.Context, s string) error {
				allowedExts := []string{".yaml", ".yml", ".opml"}
				ext := filepath.Ext(s)

				for _, allowedExt := range allowedExts {
					if ext == allowedExt {
						return nil
					}
				}

				log.Printf(`invalid file extension: "%s"`, ext)
				log.Printf("allowed extensions: %v", allowedExts)

				return errors.New("invalid file extension")
			},
		},
	}
}
