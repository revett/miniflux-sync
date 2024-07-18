package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/api"
	"github.com/revett/miniflux-sync/config"
)

func dump(cfg *config.GlobalFlags, flags *config.DumpFlags) error {
	client, err := api.Client(cfg)
	if err != nil {
		return errors.Wrap(err, "creating miniflux client")
	}

	log.Println("exporting data from miniflux")
	dat, err := client.Export()
	if err != nil {
		return errors.Wrap(err, "getting export data")
	}

	// e.g. "miniflux-sync-remote-20240718105851_opml.xml"
	filename := fmt.Sprintf("./miniflux-sync-remote-%s_opml.xml", time.Now().Format("20060102150405"))
	if flags.Path != "" {
		log.Printf(`using export path from "--path": "%s"`, flags.Path)
		filename = flags.Path
	}

	log.Println("writing export data to file")
	if err := os.WriteFile(filename, dat, 0o600); err != nil { //nolint:mnd
		return errors.Wrap(err, "writing export data to file")
	}

	log.Println(filename)
	return nil
}
