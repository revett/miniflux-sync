package diff

// CalculateDiff calculates the differences between the local and remote state and returns the
// actions to be performed.
func CalculateDiff(local *State, remote *State) ([]Action, error) {
	// TODO: Is this the correct order to ensure that we don't get strange data integrity issues?
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

	// TODO: Implement diff logic.

	return actions, nil
}
