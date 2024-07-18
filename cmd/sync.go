package cmd

import (
	"log"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/api"
	"github.com/revett/miniflux-sync/config"
)

// Sync is the entry point for the sync command in the CLI.
func Sync(cfg *config.Config) error {
	client, err := api.Client(cfg)
	if err != nil {
		return errors.Wrap(err, "creating miniflux client")
	}

	feedsByCategory, err := api.GetFeedsByCategories(client)
	if err != nil {
		return errors.Wrap(err, "getting feeds by category")
	}

	log.Printf("feeds: %d\n", len(feedsByCategory.Feeds()))
	log.Printf("categories: %d\n", len(feedsByCategory))

	return nil
}
