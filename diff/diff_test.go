package diff_test

import (
	"testing"

	"github.com/revett/miniflux-sync/diff"
	"github.com/stretchr/testify/require"
)

func TestCalculateDiff(t *testing.T) {
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
				"Music": {
					"https://music.com/feed",
				},
			},
			remote: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
				},
				"Music": {
					"https://music.com/feed",
				},
			},
			expected: []diff.Action{},
		},

		"CreateFeed": {
			local: map[string][]string{
				"Tech": {
					"https://tech.com/feed",
				},
			},
			remote: map[string][]string{
				"Music": {
					"https://music.com/feed",
				},
			},
			expected: []diff.Action{
				{
					Type:          diff.CreateFeed,
					FeedURL:       "https://tech.com/feed",
					CategoryTitle: "Tech",
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
					Type:          diff.CreateFeed,
					FeedURL:       "https://tech.com/feed",
					CategoryTitle: "Tech",
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
