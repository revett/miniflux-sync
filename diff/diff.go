package diff

import "sort"

// CalculateDiff calculates the differences between the local and remote state and returns the
// actions to be performed.
func CalculateDiff(local *State, remote *State) ([]Action, error) { //nolint:cyclop
	actions := []Action{}

	// Iterate over remote feeds and check if they exist in the local feeds.
	for categoryTitle, feedURLs := range remote.FeedURLsByCategoryTitle {
		for _, feedURL := range feedURLs {
			if !local.FeedExists(feedURL, categoryTitle) {
				actions = append(actions, Action{
					Type:          DeleteFeed,
					CategoryTitle: categoryTitle,
					FeedURL:       feedURL,
				})
			}
		}
	}

	// Iterate over remote categories and check if they exist in the local categories.
	for categoryTitle := range remote.FeedURLsByCategoryTitle {
		if !local.CategoryExists(categoryTitle) {
			actions = append(actions, Action{
				Type:          DeleteCategory,
				CategoryTitle: categoryTitle,
			})
		}
	}

	// Iterate over local categories and check if they exist in the remote categories.
	for categoryTitle := range local.FeedURLsByCategoryTitle {
		if !remote.CategoryExists(categoryTitle) {
			actions = append(actions, Action{
				Type:          CreateCategory,
				CategoryTitle: categoryTitle,
			})
		}
	}

	// Iterate over local feeds and check if they exist in the remote feeds.
	for categoryTitle, feedURLs := range local.FeedURLsByCategoryTitle {
		for _, feedURL := range feedURLs {
			if !remote.FeedExists(feedURL, categoryTitle) {
				actions = append(actions, Action{
					Type:          CreateFeed,
					CategoryTitle: categoryTitle,
					FeedURL:       feedURL,
				})
			}
		}
	}

	sort.Sort(ActionSorter(actions))

	return actions, nil
}
