package api

import (
	"log"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/config"
	miniflux "miniflux.app/v2/client"
)

// Client creates a new Miniflux API client, whilst checking the health of the Miniflux instance.
func Client(cfg *config.GlobalFlags) (*miniflux.Client, error) {
	log.Println("connecting to miniflux instance")
	client := miniflux.New(cfg.Endpoint, cfg.APIKey)

	log.Println("checking health of miniflux instance")
	if err := client.Healthcheck(); err != nil {
		return nil, errors.Wrap(err, "checking health of miniflux instance")
	}

	return client, nil
}
