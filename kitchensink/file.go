package kitchensink

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/revett/miniflux-sync/log"
)

// ValidateFileExtension checks if the file extension is in the list of allowed extensions.
func ValidateFileExtension(ctx context.Context, s string, allowedExts []string) error {
	ext := filepath.Ext(s)

	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			return nil
		}
	}

	log.Info(ctx, "invalid file extension", log.Metadata{
		"extension": ext,
	})
	log.Info(ctx, "allowed extensions", log.Metadata{
		"extensions": allowedExts,
	})

	return errors.New("invalid file extension")
}
