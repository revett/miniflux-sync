package config

import (
	"github.com/urfave/cli/v2"
)

// Config holds the configuration for the CLI.
type Config struct {
	APIKey         string
	Endpoint       string
	ExportFilename string
	Version        string
}

// New is a convenience function for creating a new Config.
func New(v string) *Config {
	return &Config{
		Version: v,
	}
}

// Flags returns the flags for the CLI.
func (c *Config) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "endpoint",
			Usage:       "Miniflux API endpoint.",
			EnvVars:     []string{"MINIFLUX_SYNC_ENDPOINT"},
			Destination: &c.Endpoint,
			Aliases:     []string{"e"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "api-key",
			Usage:       "Miniflux API key.",
			EnvVars:     []string{"MINIFLUX_SYNC_API_KEY"},
			Destination: &c.APIKey,
			Aliases:     []string{"a"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "export-filename",
			Usage:       "Filename for exported data.",
			EnvVars:     []string{"MINIFLUX_SYNC_EXPORT_FILENAME"},
			Destination: &c.ExportFilename,
			Aliases:     []string{"ef"},
		},
	}
}
