package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/api"
	"github.com/revett/miniflux-sync/config"
	"github.com/revett/miniflux-sync/log"
	"gopkg.in/yaml.v2"
	miniflux "miniflux.app/v2/client"
)

func dump(ctx context.Context, flags *config.DumpFlags, client *miniflux.Client) error {
	log.Info(ctx, "exporting data from miniflux")

	feeds, _, err := api.FetchData(ctx, client)
	if err != nil {
		return errors.Wrap(err, "fetching data")
	}

	remoteState, err := api.GenerateDiffState(feeds)
	if err != nil {
		return errors.Wrap(err, "generating remote state")
	}

	// e.g. "miniflux-sync-remote-20240718105851_opml.xml"
	filename := fmt.Sprintf("./miniflux-sync-remote-%s.yml", time.Now().Format("20060102150405"))
	if flags.Path != "" {
		log.Info(ctx, `using export path from "--path"`, log.Metadata{
			"path": flags.Path,
		})
		filename = flags.Path
	}

	log.Info(ctx, "writing export data to file")

	dat, err := yaml.Marshal(remoteState.FeedURLsByCategoryTitle)
	if err != nil {
		return errors.Wrap(err, "marshalling remote state to yaml")
	}

	if err := os.WriteFile(filename, dat, 0o600); err != nil { //nolint:mnd
		return errors.Wrap(err, "writing export data to file")
	}

	log.Info(ctx, filename)
	return nil
}
