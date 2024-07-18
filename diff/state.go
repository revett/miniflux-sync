package diff

// State represents either the current remote state of Miniflux, or the intended local state of
// Miniflux.
type State struct {
	FeedURLsByCategoryTitle map[string][]string
}

// CategoryTitles returns a list of all category titles in the state.
func (s State) CategoryTitles() []string {
	categorySet := map[string]struct{}{}

	for categoryTitle := range s.FeedURLsByCategoryTitle {
		categorySet[categoryTitle] = struct{}{}
	}

	categoryTitles := make([]string, 0, len(categorySet))
	for categoryTitle := range categorySet {
		categoryTitles = append(categoryTitles, categoryTitle)
	}

	return categoryTitles
}

// FeedExists checks if a feed URL exists in a specific category.
func (s State) FeedExists(feedURL string, categoryTitle string) bool {
	feedURLs, exists := s.FeedURLsByCategoryTitle[categoryTitle]
	if !exists {
		return false
	}

	for _, url := range feedURLs {
		if url == feedURL {
			return true
		}
	}

	return false
}

// FeedURLs returns a list of all feed URLs in the state.
func (s State) FeedURLs() []string {
	feedURLs := []string{}

	for _, urls := range s.FeedURLsByCategoryTitle {
		feedURLs = append(feedURLs, urls...)
	}

	return feedURLs
}
