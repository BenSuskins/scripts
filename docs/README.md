# Scripts

> A terminal app for running commands across multiple Git repositories.

## Overview

Scripts provides an interactive terminal menu for running configured commands across Git repositories. It shows progress and command output as each repository is processed.

## Features

- Run commands across all git repos in a directory
- Interactive menu with filtering (press `/` to search)
- Sequential execution with real-time progress
- Configurable commands via YAML

## Getting Started

### Prerequisites

- Go

### Installation

Run this command from the repository root:

```shell
go install ./cmd/scripts
```

## Usage

Run from a directory containing git repositories:

```shell
scripts
```

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `↑/↓` | Navigate menu |
| `Enter` | Run selected command |
| `/` | Filter commands |
| `H` | Toggle help |
| `S` | Toggle status bar |
| `P` | Toggle pagination |
| `T` | Toggle title |
| `q` | Quit |

## Configuration

Config is loaded from `~/.config/scripts/config.yaml` or `./scripts.yaml`.

```yaml
commands:
  - name: "Branch"
    description: "Show current branch for all repos"
    command: "git rev-parse --abbrev-ref HEAD"
    type: "git_dirs"
    show_output: true

  - name: "Pull Main"
    description: "Pull main for all repos"
    command: "git pull origin main"
    type: "git_dirs"
    show_output: false
```

### Command Types

- `git_dirs` - Runs in each git repository found in current directory
- `single` - Runs once in current working directory

### Show Output

- `true` - Display command output inline (e.g., branch names)
- `false` - Only show success/failure icons

## Commands\n\nRun these commands from the repository root.\n\n## Commands

Run these commands from the repository root.

## Development

```shell
make build    # Build binary
make run      # Run the app
make test     # Run tests
make fmt      # Format code
make vet      # Run go vet
make tidy     # Tidy dependencies
make clean    # Remove binary
```


## Security

See [SECURITY.md](SECURITY.md) for vulnerability reporting.


## License

No license file is present. All rights remain with the author.
