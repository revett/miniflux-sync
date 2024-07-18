package api

import (
	"log"

	"github.com/pkg/errors"
	miniflux "miniflux.app/v2/client"
)

// WithoutCategoryID is the ID used for feeds without a category.
const WithoutCategoryID = int64(-1234)

// FeedsByCategories is a map of feeds grouped by category.
type FeedsByCategories map[int64][]*miniflux.Feed

// Feeds returns all feeds in the map.
func (f FeedsByCategories) Feeds() []*miniflux.Feed {
	var feeds []*miniflux.Feed

	for _, fs := range f {
		feeds = append(feeds, fs...)
	}

	return feeds
}

// GetFeedsByCategories returns a map of feeds grouped by category, including support for feeds
// without a category which use the WithoutCategoryID constant as their category ID in the map.
func GetFeedsByCategories(client *miniflux.Client) (FeedsByCategories, error) {
	log.Println("fetching feeds")

	feeds, err := client.Feeds()
	if err != nil {
		return nil, errors.Wrap(err, "getting feeds")
	}

	feedsByCategories := make(map[int64][]*miniflux.Feed)

	for _, feed := range feeds {
		id := WithoutCategoryID
		if feed.Category != nil {
			id = feed.Category.ID
		}

		feedsByCategories[id] = append(feedsByCategories[id], feed)
	}

	return feedsByCategories, nil
}
