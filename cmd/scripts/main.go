package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"suskins/scripts/internal/config"
	"suskins/scripts/internal/ui"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(ui.NewModel(cfg), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
