package diff_test

import (
	"sort"
	"testing"

	"github.com/revett/miniflux-sync/diff"
	"github.com/stretchr/testify/require"
)

func TestActionSorter(t *testing.T) { //nolint:funlen
	t.Parallel()

	tests := map[string]struct {
		input    []diff.Action
		expected []diff.Action
	}{
		"AlreadySorted": {
			input: []diff.Action{
				{
					Type:    diff.DeleteFeed,
					FeedURL: "https://a.com/feed",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "CategoryA",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "CategoryB",
				},
				{
					Type:    diff.CreateFeed,
					FeedURL: "https://b.com/feed",
				},
			},
			expected: []diff.Action{
				{
					Type:    diff.DeleteFeed,
					FeedURL: "https://a.com/feed",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "CategoryA",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "CategoryB",
				},
				{
					Type:    diff.CreateFeed,
					FeedURL: "https://b.com/feed",
				},
			},
		},

		"Unsorted": {
			input: []diff.Action{
				{
					Type:    diff.CreateFeed,
					FeedURL: "https://b.com/feed",
				},
				{
					Type:    diff.DeleteFeed,
					FeedURL: "https://a.com/feed",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "CategoryB",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "CategoryA",
				},
			},
			expected: []diff.Action{
				{
					Type:    diff.DeleteFeed,
					FeedURL: "https://a.com/feed",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "CategoryA",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "CategoryB",
				},
				{
					Type:    diff.CreateFeed,
					FeedURL: "https://b.com/feed",
				},
			},
		},

		"ComplexUnsorted": {
			input: []diff.Action{
				{
					Type:    diff.CreateFeed,
					FeedURL: "https://c.com/feed",
				},
				{
					Type:    diff.DeleteFeed,
					FeedURL: "https://a.com/feed",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "CategoryB",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "CategoryA",
				},
				{
					Type:    diff.CreateFeed,
					FeedURL: "https://b.com/feed",
				},
				{
					Type:    diff.DeleteFeed,
					FeedURL: "https://b.com/feed",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "CategoryA",
				},
			},
			expected: []diff.Action{
				{
					Type:    diff.DeleteFeed,
					FeedURL: "https://a.com/feed",
				},
				{
					Type:    diff.DeleteFeed,
					FeedURL: "https://b.com/feed",
				},
				{
					Type:          diff.DeleteCategory,
					CategoryTitle: "CategoryA",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "CategoryA",
				},
				{
					Type:          diff.CreateCategory,
					CategoryTitle: "CategoryB",
				},
				{
					Type:    diff.CreateFeed,
					FeedURL: "https://b.com/feed",
				},
				{
					Type:    diff.CreateFeed,
					FeedURL: "https://c.com/feed",
				},
			},
		},
	}

	for name, testCase := range tests {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			sort.Sort(diff.ActionSorter(tc.input))
			require.Equal(t, tc.expected, tc.input)
		})
	}
}
