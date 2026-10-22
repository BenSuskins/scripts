# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
go run ./cmd/scripts      # Run the application
go build ./cmd/scripts    # Build binary
go mod tidy               # Tidy dependencies
go fmt ./...              # Format code
```

## Architecture

This is a BubbleTea-based terminal UI for running commands across multiple git repositories.

### Core Flow
1. `cmd/scripts/main.go` - Entry point, loads config and starts BubbleTea program
2. `internal/config/` - YAML config loading from `~/.config/scripts/config.yaml` or `./scripts.yaml`
3. `internal/ui/model.go` - Main BubbleTea model with state machine (StateMenu → StateExecuting → StateResults)
4. `internal/core/` - Git directory scanning and command execution

### Command Types
- `git_dirs` - Runs command in each git repository found in current directory (sequential with spinner)
- `single` - Runs command once in current working directory

### Config Schema
```yaml
commands:
  - name: "Command Name"
    description: "Shown in menu"
    command: "shell command"
    type: git_dirs | single
    show_output: true | false  # false = only show ✓/✗, errors always shown
```

### Key Files
- `internal/ui/model.go` - BubbleTea Update/View logic, handles sequential command execution
- `internal/ui/menu.go` - MenuItem struct wrapping config commands for bubbles/list
- `internal/core/scanner.go` - `FindGitDirectories()` walks cwd for .git folders
- `internal/core/command.go` - `RunCommand()` executes shell commands via `sh -c`
