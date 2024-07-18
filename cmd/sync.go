package cmd

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/api"
	"github.com/revett/miniflux-sync/config"
)

func sync(cfg *config.GlobalFlags, flags *config.SyncFlags) error {
	log.Println("reading data from file")
	log.Println(flags.Path)

	_, err := os.ReadFile(flags.Path)
	if err != nil {
		return errors.Wrap(err, "reading data from file")
	}

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
