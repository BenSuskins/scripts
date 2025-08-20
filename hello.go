package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	progress progress.Model
	done     bool
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.done {
			return m, nil
		}

		cmd := m.progress.IncrPercent(0.05)
		if m.progress.Percent() >= 1.0 {
			m.done = true
			return m, tea.Quit
		}

		return m, tea.Batch(cmd, tick())

	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	var newModel tea.Model
	newModel, cmd = m.progress.Update(msg)
	m.progress = newModel.(progress.Model)

	return m, cmd
}

func (m model) View() string {
	if m.done {
		return "✅ Done!\n"
	}
	return fmt.Sprintf("\nProgress:\n%s\n\nPress q to quit.\n", m.progress.View())
}

func tick() tea.Cmd {
	return tea.Tick(time.Second/4, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	p := tea.NewProgram(model{
		progress: progress.New(progress.WithDefaultGradient()),
	})
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
