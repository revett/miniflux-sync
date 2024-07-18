package api

import (
	"log"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/diff"
	miniflux "miniflux.app/v2/client"
)

// Update performs a series of actions on the Miniflux instance.
func Update( //nolint:cyclop
	client *miniflux.Client,
	actions []diff.Action,
	feeds []*miniflux.Feed,
	categories []*miniflux.Category,
) error {
	log.Println("performing actions")

	for _, action := range actions {
		switch action.Type {
		case diff.CreateCategory:
			log.Printf(`creating category: "%s"`, action.CategoryTitle)
			if _, err := client.CreateCategory(action.CategoryTitle); err != nil {
				return errors.Wrap(err, "creating category")
			}

		case diff.CreateFeed:
			log.Printf(`creating feed: "%s / %s"`, action.CategoryTitle, action.FeedURL)
			categoryID, err := findCategoryIDByTitle(action.CategoryTitle, categories)
			if err != nil {
				return errors.Wrap(err, "finding category id")
			}

			req := miniflux.FeedCreationRequest{
				FeedURL:    action.FeedURL,
				CategoryID: categoryID,
			}

			if _, err := client.CreateFeed(&req); err != nil {
				return errors.Wrap(err, "creating feed")
			}

		case diff.DeleteCategory:
			log.Printf(`deleting category: "%s"`, action.CategoryTitle)
			categoryID, err := findCategoryIDByTitle(action.CategoryTitle, categories)
			if err != nil {
				return errors.Wrap(err, "finding category id")
			}

			if err := client.DeleteCategory(categoryID); err != nil {
				return errors.Wrap(err, "deleting category")
			}

		case diff.DeleteFeed:
			log.Printf(`deleting feed: "%s / %s"`, action.CategoryTitle, action.FeedURL)
			feedID, err := findFeedIDByURL(action.FeedURL, feeds)
			if err != nil {
				return errors.Wrap(err, "finding feed id")
			}

			if err := client.DeleteFeed(feedID); err != nil {
				return errors.Wrap(err, "deleting feed")
			}

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