package api

import (
	"log"

	"github.com/pkg/errors"
	miniflux "miniflux.app/v2/client"
)

func FetchData(client *miniflux.Client) ([]*miniflux.Feed, []*miniflux.Category, error) {
	log.Println("fetching feeds")

	feeds, err := client.Feeds()
	if err != nil {
		return nil, nil, errors.Wrap(err, "fetching feeds")
	}

	log.Println("fetching categories")

	categories, err := client.Categories()
	if err != nil {
		return nil, nil, errors.Wrap(err, "fetching categories")
	}

	return feeds, categories, nil
}
