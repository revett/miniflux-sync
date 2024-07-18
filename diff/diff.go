package diff

// CalculateDiff calculates the differences between the local and remote state and returns the
// actions to be performed.
func CalculateDiff(local *State, remote *State) ([]Action, error) {
	// TODO: Is this the correct order to ensure that we don't get strange data integrity issues?
	actions := []Action{}

	// Create a map for quick lookup of remote feeds.
	remoteFeeds := make(map[string]string)
	for categoryTitle, feedURLs := range remote.FeedURLsByCategoryTitle {
		for _, feedURL := range feedURLs {
			remoteFeeds[feedURL] = categoryTitle
		}
	}

	// Iterate over local feeds and check if they exist in the remote feeds.
	for categoryTitle, feedURLs := range local.FeedURLsByCategoryTitle {
		for _, feedURL := range feedURLs {
			if remoteCategory, exists := remoteFeeds[feedURL]; !exists || remoteCategory != categoryTitle {
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
