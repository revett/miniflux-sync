#!/bin/bash

# Extract the Go version from go.mod and print the major and minor parts only.
awk '/^go / {print $2}' go.mod | cut -d. -f1,2
