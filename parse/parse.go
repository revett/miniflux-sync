package parse

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/diff"
	"gopkg.in/yaml.v2"
)

// Parse reads a YAML file to a diff.State struct.
func Parse(path string) (*diff.State, error) {
	log.Println("reading data from yaml file")
	log.Println(path)

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

	// TODO: Add validation for duplicate feed URLs or categories.

	return &state, nil
}
