package diff_test

import (
	"testing"

	"github.com/revett/miniflux-sync/diff"
	"github.com/stretchr/testify/require"
)

// TODO: Add test case for renaming a category with feeds in it.

func TestCalculateDiff(t *testing.T) { //nolint:funlen
	t.Parallel()

	tests := map[string]struct {
		local    map[string][]string
		remote   map[string][]string
		expected []diff.Action
	}{
		"NoAction": {
			local: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
				},
			},
			remote: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
				},
			},
			expected: []diff.Action{},
		},

		"CreateFeed": {
			local: map[string][]string{
				"General": {
					"https://tech.com/feed",
					"https://music.com/feed",
				},
			},
			remote: map[string][]string{
				"General": {
					"https://tech.com/feed",
				},
			},
			expected: []diff.Action{
				{
					Type:          diff.CreateFeed,
					FeedURL:       "https://music.com/feed",
					CategoryTitle: "General",
				},
			},
		},

		"CreateFeedInDifferentCategory": {
			local: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
				},
			},
			remote: map[string][]string{
				"General": {
					"https://tech.com/feed",
				},
			},
			expected: []diff.Action{
				{
					Type:          diff.DeleteFeed,
					FeedURL:       "https://tech.com/feed",
					CategoryTitle: "General",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "General",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "Tech",
				},
				{
					Type:          diff.CreateFeed,
					FeedURL:       "https://tech.com/feed",
					CategoryTitle: "Tech",
				},
			},
		},

		"DeleteFeed": {
			local: map[string][]string{},
			remote: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
				},
			},
			expected: []diff.Action{
				{
					Type:          diff.DeleteFeed,
					FeedURL:       "https://tech.com/feed",
					CategoryTitle: "Tech",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "Tech",
				},
			},
		},

		"CreateCategoryAndFeed": {
			local: map[string][]string{
				"Music": {
					"https://music.com/feed",
				},
				"Tech": {
					"https://tech.com/feed",
				},
			},
			remote: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
				},
			},
			expected: []diff.Action{
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "Music",
				},
				{
					Type:          diff.CreateFeed,
					FeedURL:       "https://music.com/feed",
					CategoryTitle: "Music",
				},
			},
		},

		"DeleteCategoryAndFeed": {
			local: map[string][]string{
				"Music": {
					"https://music.com/feed",
				},
			},
			remote: map[string][]string{
				"Music": {
					"https://music.com/feed",
				},
				"Tech": {
					"https://tech.com/feed",
				},
			},
			expected: []diff.Action{
				{
					Type:          diff.DeleteFeed,
					FeedURL:       "https://tech.com/feed",
					CategoryTitle: "Tech",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "Tech",
				},
			},
		},

		"CreateMultipleCategoriesAndFeeds": {
			local: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
					"https://newtech.com/feed",
				},
				"Music": {
					"https://music.com/feed",
				},
				"News": {
					"https://news.com/feed",
				},
			},
			remote: map[string][]string{},
			expected: []diff.Action{
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "Music",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "News",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "Tech",
				},
				{
					Type:          diff.CreateFeed,
					FeedURL:       "https://music.com/feed",
					CategoryTitle: "Music",
				},
				{
					Type:          diff.CreateFeed,
					FeedURL:       "https://news.com/feed",
					CategoryTitle: "News",
				},
				{
					Type:          diff.CreateFeed,
					FeedURL:       "https://newtech.com/feed",
					CategoryTitle: "Tech",
				},
				{
					Type:          diff.CreateFeed,
					FeedURL:       "https://tech.com/feed",
					CategoryTitle: "Tech",
				},
			},
		},

		"DeleteMultipleFeedsInDifferentCategories": {
			local: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
				},
			},
			remote: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
					"https://oldtech.com/feed",
				},
				"Music": {
					"https://music.com/feed",
				},
			},
			expected: []diff.Action{
				{
					Type:          diff.DeleteFeed,
					FeedURL:       "https://music.com/feed",
					CategoryTitle: "Music",
				},
				{
					Type:          diff.DeleteFeed,
					FeedURL:       "https://oldtech.com/feed",
					CategoryTitle: "Tech",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "Music",
				},
			},
		},
	}

	for name, testCase := range tests {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			local := &diff.State{
				FeedURLsByCategoryTitle: tc.local,
			}
			remote := &diff.State{
				FeedURLsByCategoryTitle: tc.remote,
			}

			diff, err := diff.CalculateDiff(local, remote)
			require.NoError(t, err)
			require.Equal(t, tc.expected, diff)
		})
	}
}
