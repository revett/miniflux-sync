package api

import (
	"context"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/log"
	miniflux "miniflux.app/v2/client"
)

// FetchData fetches feeds and categories from the Miniflux instance.
func FetchData(
	ctx context.Context, client *miniflux.Client,
) ([]*miniflux.Feed, []*miniflux.Category, error) {
	log.Info(ctx, "fetching feeds")

	feeds, err := client.Feeds()
	if err != nil {
		return nil, nil, errors.Wrap(err, "fetching feeds")
	}

	log.Info(ctx, "fetching categories")

	categories, err := client.Categories()
	if err != nil {
		return nil, nil, errors.Wrap(err, "fetching categories")
	}

	return feeds, categories, nil
}
