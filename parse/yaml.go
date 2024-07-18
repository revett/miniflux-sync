package parse

import (
	"fmt"
	"net/url"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// LoadYAML reads a YAML file, validates the structure, and returns a map of categories and feed
// URLs.
func LoadYAML(path string) (map[string][]string, error) {
	data, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		return nil, errors.Wrap(err, "reading data from file")
	}

	feeds := make(map[string][]string)

	if err := yaml.Unmarshal(data, &feeds); err != nil {
		return nil, errors.Wrap(err, "unmarshalling data")
	}

	for category, urls := range feeds {
		if category == "" {
			return nil, errors.New("empty category name found in yaml file")
		}

		for _, u := range urls {
			if _, err := url.ParseRequestURI(u); err != nil {
				msg := fmt.Sprintf(`invalid URL "%s" in category "%s"`, u, category)
				return nil, errors.Wrap(err, msg)
			}
		}
	}

	return feeds, nil
}
