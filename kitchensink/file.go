package kitchensink

import (
	"errors"
	"log"
	"path/filepath"
)

// ValidateFileExtension checks if the file extension is in the list of allowed extensions.
func ValidateFileExtension(s string, allowedExts []string) error {
	ext := filepath.Ext(s)

	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			return nil
		}
	}

	log.Printf(`invalid file extension: "%s"`, ext)
	log.Printf("allowed extensions: %v", allowedExts)

	return errors.New("invalid file extension")
}
