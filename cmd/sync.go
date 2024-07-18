package cmd

import (
	"log"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/api"
	"github.com/revett/miniflux-sync/config"
	"github.com/revett/miniflux-sync/parse"
)

func sync(cfg *config.GlobalFlags, flags *config.SyncFlags) error {
	log.Println("reading data from file")
	log.Println(flags.Path)

	// TODO: Add dry run support.

	switch filepath.Ext(flags.Path) {
	case ".yaml", ".yml":
		log.Println("importing from yaml file")

		_, err := parse.LoadYAML(flags.Path)
		if err != nil {
			return errors.Wrap(err, "loading data from yaml file")
		}

		// TODO: Implement logic for YAML.

	case ".opml":
		log.Println("importing from opml file")
		return errors.New("opml file format not implemented")

		// TODO: Implement logic for OPML.

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

	// TODO: Implement diff logic.
	// TODO: Implement update logic.

	return nil
}
