# Scripts

A terminal UI for running commands across multiple git repositories.

## Features

- Run commands across all git repos in a directory
- Interactive menu with filtering (press `/` to search)
- Sequential execution with real-time progress
- Configurable commands via YAML

## Installation

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
