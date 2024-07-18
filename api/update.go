package api

import (
	"context"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/diff"
	"github.com/revett/miniflux-sync/log"
	miniflux "miniflux.app/v2/client"
)

// Update performs a series of actions on the Miniflux instance.
func Update( //nolint:cyclop,funlen
	ctx context.Context,
	client *miniflux.Client,
	actions []diff.Action,
	feeds []*miniflux.Feed,
	categories []*miniflux.Category,
) error {
	log.Info(ctx, "performing actions")

	for _, action := range actions {
		switch action.Type {
		case diff.CreateCategory:
			log.Info(ctx, "creating category", log.Metadata{
				"title": action.CategoryTitle,
			})

			category, err := client.CreateCategory(action.CategoryTitle)
			if err != nil {
				return errors.Wrap(err, "creating category")
			}

			categories = append(categories, category)

		case diff.CreateFeed:
			log.Info(ctx, "creating feed", log.Metadata{
				"category": action.CategoryTitle,
				"url":      action.FeedURL,
			})

			categoryID, err := findCategoryIDByTitle(action.CategoryTitle, categories)
			if err != nil {
				return errors.Wrap(err, "finding category id")
			}

			req := miniflux.FeedCreationRequest{
				FeedURL:    action.FeedURL,
				CategoryID: categoryID,
			}

			feedID, err := client.CreateFeed(&req)
			if err != nil {
				return errors.Wrap(err, "creating feed")
			}

			feed, err := client.Feed(feedID)
			if err != nil {
				return errors.Wrap(err, "fetching feed")
			}

			feeds = append(feeds, feed)

		case diff.DeleteCategory:
			log.Info(ctx, "deleting category", log.Metadata{
				"title": action.CategoryTitle,
			})

			categoryID, err := findCategoryIDByTitle(action.CategoryTitle, categories)
			if err != nil {
				return errors.Wrap(err, "finding category id")
			}

			if err := client.DeleteCategory(categoryID); err != nil {
				return errors.Wrap(err, "deleting category")
			}

			categories = removeCategoryByID(categoryID, categories)

		case diff.DeleteFeed:
			log.Info(ctx, "deleting feed", log.Metadata{
				"category": action.CategoryTitle,
				"url":      action.FeedURL,
			})

			feedID, err := findFeedIDByURL(action.FeedURL, feeds)
			if err != nil {
				return errors.Wrap(err, "finding feed id")
			}

			if err := client.DeleteFeed(feedID); err != nil {
				return errors.Wrap(err, "deleting feed")
			}

			feeds = removeFeedByID(feedID, feeds)

		default:
			return errors.Errorf(`unknown action type: "%s"`, action.Type)
		}
	}

	return nil
}

func findCategoryIDByTitle(title string, categories []*miniflux.Category) (int64, error) {
	for _, category := range categories {
		if category.Title == title {
			return category.ID, nil
		}
	}

	return 0, errors.Errorf(`category not found: "%s"`, title)
}

func findFeedIDByURL(url string, feeds []*miniflux.Feed) (int64, error) {
	for _, feed := range feeds {
		if feed.FeedURL == url {
			return feed.ID, nil
		}
	}

	return 0, errors.Errorf(`feed not found: "%s"`, url)
}

func removeCategoryByID(id int64, categories []*miniflux.Category) []*miniflux.Category {
	for i, category := range categories {
		if category.ID == id {
			return append(categories[:i], categories[i+1:]...)
		}
	}

	return categories
}

func removeFeedByID(id int64, feeds []*miniflux.Feed) []*miniflux.Feed {
	for i, feed := range feeds {
		if feed.ID == id {
			return append(feeds[:i], feeds[i+1:]...)
		}
	}

	return feeds
}
