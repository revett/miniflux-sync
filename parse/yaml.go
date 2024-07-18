package parse

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// LoadYAML reads a YAML file to a LocalConfig.
func LoadYAML(path string) (*LocalConfig, error) {
	log.Println("reading data from yaml file")
	log.Println(path)

	data, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		return nil, errors.Wrap(err, "reading data from file")
	}

	localConfig := LocalConfig{
		FeedsByCategory: make(map[string][]string),
	}

	if err := yaml.Unmarshal(data, &localConfig.FeedsByCategory); err != nil {
		return nil, errors.Wrap(err, "unmarshalling data")
	}

	return &localConfig, nil
}
