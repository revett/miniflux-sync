# miniflux-sync

![GitHub Release](https://img.shields.io/github/v/release/revett/miniflux-sync?label=Release)
![GitHub branch check runs](https://img.shields.io/github/check-runs/revett/miniflux-sync/main?label=Checks)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/revett/miniflux-sync)
![GitHub License](https://img.shields.io/github/license/revett/miniflux-sync?label=License)

Manage and sync your [Miniflux](https://github.com/miniflux/v2) feeds with YAML.

![YAML config](./.resources/config.png)
![Logs](./.resources/logs.png)

## Install

- Download the [latest release](https://github.com/revett/miniflux-sync/releases)
- For macOS, follow these [steps](https://support.apple.com/en-il/guide/mac-help/mchleab3a043/mac)

## Usage

Configure the CLI to use and authenticate with your Miniflux instance:

```bash
# Use environment variables
MINIFLUX_SYNC_ENDPOINT=... MINIFLUX_SYNC_API_KEY=... miniflux-sync -h

# Or via CLI flags
miniflux-sync --endpoint="..." --api-key="..." -h
```

Then run the CLI:

```bash
# Help
miniflux-sync -h

# View changes via dry run
miniflux-sync sync --path ./feeds.yml --dry-run

# Sync changes
miniflux-sync sync --path ./feeds.yml

# Export remote state
miniflux-sync dump
```

## Contributing

Contributions, issues and feature requests are very welcome.

```bash
# Running tests
go test -cover ./...

# Bump VERSION, and run script
GITHUB_TOKEN="..." ./scripts/release.sh
```
