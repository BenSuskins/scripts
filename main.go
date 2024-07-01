package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
	action      func()
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				i.action()
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	items := []list.Item{
		item{title: "Pull All", desc: "Pull's all git repo's in current directory on their current branch", action: pullAll},
		item{title: "Pull All - Master", desc: "Pull's all git repo's in current directory on the master branch", action: pullAll},
		item{title: "Pull All - Develop", desc: "Pull's all git repo's in current directory on the develop branch", action: pullAll},
		item{title: "Update", desc: "Updates system and tools", action: updateSystem},
		item{title: "Cleanup", desc: "Cleans up system", action: cleanup},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Scripts"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// Function to pull all (example implementation)
func pullAll() {
	fmt.Println("Executing: Pulling All")
	// Your Go code to pull all
}

// Function to update Docker Compose (example implementation)
func updateDockerCompose() {
	fmt.Println("Executing: Updating Docker Compose")
	// Your Go code to update Docker Compose
}

// Function to update the system (example implementation)
func updateSystem() {
	fmt.Println("Executing: Updating System")
	// Your Go code to update the system
}

// Function to cleanup (example implementation)
func cleanup() {
	fmt.Println("Executing: Cleanup")
	// Your Go code to cleanup
}