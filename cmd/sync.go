package cmd

import (
	"log"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/revett/miniflux-sync/api"
	"github.com/revett/miniflux-sync/config"
	"github.com/revett/miniflux-sync/diff"
	"github.com/revett/miniflux-sync/parse"
	miniflux "miniflux.app/v2/client"
)

func sync(flags *config.SyncFlags, client *miniflux.Client) error { //nolint:cyclop
	var localState *diff.State
	var err error

	switch filepath.Ext(flags.Path) {
	case ".yaml", ".yml":
		localState, err = parse.Parse(flags.Path)
		if err != nil {
			return errors.Wrap(err, "loading data from yaml file")
		}

	default:
		return errors.New("invalid file extension") // Should never happen, as we validate flag before.
	}

	log.Printf("local feeds: %d", len((localState.FeedURLs())))
	log.Printf("local categories: %d", len(localState.CategoryTitles()))

	feeds, categories, err := api.FetchData(client)
	if err != nil {
		return errors.Wrap(err, "fetching data")
	}

	remoteState, err := api.GenerateDiffState(feeds)
	if err != nil {
		return errors.Wrap(err, "generating remote state")
	}

	log.Printf("remote feeds: %d", len(remoteState.FeedURLs()))
	log.Printf("remote categories: %d", len(remoteState.CategoryTitles()))

	actions, err := diff.CalculateDiff(localState, remoteState)
	if err != nil {
		return errors.Wrap(err, "calculating diff")
	}

	if len(actions) == 0 {
		log.Printf("no actions to perform")
		return nil
	}

	log.Printf("%d actions to perform:", len(actions))
	for _, action := range actions {
		log.Printf(`%s: "%s / %s"`, action.Type, action.CategoryTitle, action.FeedURL)
	}

	if flags.DryRun {
		log.Println("dry run complete")
		return nil
	}

	if err := api.Update(client, actions, feeds, categories); err != nil {
		return errors.Wrap(err, "performing actions")
	}

	return nil
}
