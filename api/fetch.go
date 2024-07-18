package api

import (
	"log"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/diff"
	miniflux "miniflux.app/v2/client"
)

// FetchState fetches all feeds from the Miniflux client and returns them as a diff.State struct.
func FetchState(client *miniflux.Client) (*diff.State, error) {
	log.Println("fetching feeds")

	feeds, err := client.Feeds()
	if err != nil {
		return nil, errors.Wrap(err, "getting feeds")
	}

	state := diff.State{
		FeedURLsByCategoryTitle: map[string][]string{},
	}

	// Populate state with values, and create category set.
	for _, feed := range feeds {
		if feed.Category == nil {
			return nil, errors.New("feed has no category")
		}
		categoryTitle := feed.Category.Title

		state.FeedURLsByCategoryTitle[categoryTitle] = append(
			state.FeedURLsByCategoryTitle[categoryTitle], feed.FeedURL,
		)
	}

	return &state, nil
}
