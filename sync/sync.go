package sync

import (
	"log"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/config"
	miniflux "miniflux.app/v2/client"
)

// Sync reads Miniflux feeds and categories from a YAML file and syncs them with a MiniFlux instance
// via the Miniflux API.
func Sync(cfg *config.Config) error {
	log.Println("syncing config")

	client := miniflux.New(cfg.Endpoint, cfg.APIKey)
	if err := client.Healthcheck(); err != nil {
		return errors.Wrap(err, "checking health of miniflux instance")
	}

	feeds, err := client.Feeds()
	if err != nil {
		return errors.Wrap(err, "getting feeds")
	}

	categoriesWithFeeds := make(map[int64][]int64)
	feedsWithoutCategory := []int64{}

	for _, feed := range feeds {
		if feed.Category == nil {
			feedsWithoutCategory = append(feedsWithoutCategory, feed.ID)
			continue
		}

		id := feed.Category.ID
		categoriesWithFeeds[id] = append(categoriesWithFeeds[id], feed.ID)
	}

	log.Printf("feeds: %d\n", len(feeds))
	log.Printf("categories: %d\n", len(categoriesWithFeeds))

	if len(feedsWithoutCategory) > 0 {
		log.Printf("feeds without category: %d\n", len(feedsWithoutCategory))
	}

	return nil
}
