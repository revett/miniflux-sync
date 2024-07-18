package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/config"
	"github.com/revett/miniflux-sync/log"
	miniflux "miniflux.app/v2/client"
)

func dump(ctx context.Context, flags *config.DumpFlags, client *miniflux.Client) error {
	log.Info(ctx, "exporting data from miniflux")
	dat, err := client.Export()
	if err != nil {
		return errors.Wrap(err, "getting export data")
	}

	// e.g. "miniflux-sync-remote-20240718105851_opml.xml"
	filename := fmt.Sprintf("./miniflux-sync-remote-%s_opml.xml", time.Now().Format("20060102150405"))
	if flags.Path != "" {
		log.Info(ctx, `using export path from "--path"`, log.Metadata{
			"path": flags.Path,
		})
		filename = flags.Path
	}

	log.Info(ctx, "writing export data to file")
	if err := os.WriteFile(filename, dat, 0o600); err != nil { //nolint:mnd
		return errors.Wrap(err, "writing export data to file")
	}

	log.Info(ctx, filename)
	return nil
}
