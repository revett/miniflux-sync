package config

import (
	"github.com/urfave/cli/v2"
)

// GlobalFlags holds the configuration for the CLI.
type GlobalFlags struct {
	APIKey   string
	Endpoint string
	Version  string
}

// New is a convenience function for creating a new Config.
func New(v string) *GlobalFlags {
	return &GlobalFlags{
		Version: v,
	}
}

// Flags returns the flags for the CLI.
func (c *GlobalFlags) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "api-key",
			Usage:       "Miniflux API key. (required)",
			EnvVars:     []string{"MINIFLUX_SYNC_API_KEY"},
			Destination: &c.APIKey,
			Aliases:     []string{"a"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "endpoint",
			Usage:       "Miniflux API endpoint. (required)",
			EnvVars:     []string{"MINIFLUX_SYNC_ENDPOINT"},
			Destination: &c.Endpoint,
			Aliases:     []string{"e"},
			Required:    true,
		},
	}
}
