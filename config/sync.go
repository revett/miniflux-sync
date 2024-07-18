package config

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/kitchensink"
	"github.com/urfave/cli/v2"
)

// SyncFlags holds the flags for the sync command.
type SyncFlags struct {
	DryRun bool
	Path   string
}

// Flags returns the flags for the sync command.
func (s *SyncFlags) Flags(ctx context.Context) []cli.Flag {
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
			Action: func(_ *cli.Context, path string) error {
				if err := kitchensink.ValidateFileExtension(
					ctx, path, []string{".yaml", ".yml"},
				); err != nil {
					return errors.Wrap(err, "validating file extension")
				}

				info, err := os.Stat(path)
				if os.IsNotExist(err) {
					return errors.Wrap(err, "file does not exist")
				}

				if info.IsDir() {
					return errors.New("file is a directory")
				}

				return nil
			},
		},
	}
}
