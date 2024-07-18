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
	var localConfig *parse.LocalConfig
	var err error

	// TODO: Add dry run support.

	switch filepath.Ext(flags.Path) {
	case ".yaml", ".yml":
		localConfig, err = parse.LoadYAML(flags.Path)
		if err != nil {
			return errors.Wrap(err, "loading data from yaml file")
		}

	default:
		return errors.New("invalid file extension") // Should never happen, as we validate flag before.
	}

	log.Printf("local feeds: %d\n", len(localConfig.FeedsByCategory))
	log.Printf("local categories: %d\n", len(localConfig.FeedsByCategory))

	client, err := api.Client(cfg)
	if err != nil {
		return errors.Wrap(err, "creating miniflux client")
	}

	feedsByCategory, err := api.GetFeedsByCategories(client)
	if err != nil {
		return errors.Wrap(err, "getting feeds by category")
	}

	// TODO: feedsByCategory and localConfig do the same thing, merge them in a new package.
	log.Printf("remote feeds: %d\n", len(feedsByCategory.Feeds()))
	log.Printf("remote categories: %d\n", len(feedsByCategory))

	// TODO: Implement diff logic.
	// TODO: Implement update logic.

	return nil
}
