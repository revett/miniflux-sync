package parse

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
)

// LocalConfig holds the configuration from a user for what feeds to import and categorise.
type LocalConfig struct {
	FeedsByCategory map[string][]string
}

// Feeds returns all feeds in the map.
func (l LocalConfig) Feeds() []string {
	var feeds []string

	for _, fs := range l.FeedsByCategory {
		feeds = append(feeds, fs...)
	}

	return feeds
}

// Validate checks the LocalConfig for any invalid data.
func (l LocalConfig) Validate() error {
	for category, urls := range l.FeedsByCategory {
		if category == "" {
			return errors.New("empty category name found in yaml file")
		}

		for _, u := range urls {
			if _, err := url.ParseRequestURI(u); err != nil {
				msg := fmt.Sprintf(`invalid URL "%s" in category "%s"`, u, category)
				return errors.Wrap(err, msg)
			}
		}
	}

	return nil
}
