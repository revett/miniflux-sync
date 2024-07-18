package parse

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/diff"
	"github.com/revett/miniflux-sync/log"
	"gopkg.in/yaml.v2"
)

// Parse reads a YAML file to a diff.State struct.
func Parse(ctx context.Context, path string) (*diff.State, error) {
	log.Info(ctx, "reading data from yaml file")
	log.Info(ctx, path)

	data, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		return nil, errors.Wrap(err, "reading data from file")
	}

	state := diff.State{
		FeedURLsByCategoryTitle: map[string][]string{},
	}

	if err := yaml.Unmarshal(data, &state.FeedURLsByCategoryTitle); err != nil {
		return nil, errors.Wrap(err, "unmarshalling data")
	}

	if err := validateDuplicateFeedURLs(&state); err != nil {
		return nil, errors.Wrap(err, "validating duplicate feed urls")
	}

	return &state, nil
}

func validateDuplicateFeedURLs(state *diff.State) error {
	feedURLSet := make(map[string]struct{})

	for _, urls := range state.FeedURLsByCategoryTitle {
		for _, url := range urls {
			if _, exists := feedURLSet[url]; exists {
				return errors.Errorf(`duplicate url found across categories: "%s"`, url)
			}

			feedURLSet[url] = struct{}{}
		}
	}

	return nil
}
