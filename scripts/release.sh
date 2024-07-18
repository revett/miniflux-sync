#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Read the version from the VERSION file.
VERSION=$(cat VERSION)
echo "version: $VERSION"

# Check if the tag already exists in the remote repository using gh CLI.
echo "checking if tag $VERSION already exists in the remote repo"
if gh release view "$VERSION" >/dev/null 2>&1; then
  echo "Tag $VERSION already exists."
  exit 1
fi

# Create the tag and push it to the remote repository.
echo "tag $VERSION does not exist, creating and pushing tag"
git tag "$VERSION"
git push origin "$VERSION"
echo "tag $VERSION created and pushed successfully"

# Run goreleaser release.
echo "running goreleaser"
goreleaser release
echo "goreleaser release completed successfully"
