package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().
	Margin(1, 2).
	Bold(true).
	PaddingTop(2).
	PaddingLeft(4)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.desc
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Have a good day!")
	}
	return docStyle.Render(m.list.View())
}

var quitTextStyle = lipgloss.NewStyle().
	Margin(1, 2).
	Bold(true).
	Foreground(lipgloss.Color("205")).
	PaddingTop(1).
	PaddingLeft(2)

func main() {
	items := []list.Item{
		item{title: "Pull All", desc: "Pull's all git repo's in current directory on their current branch"},
		item{title: "Pull All - Master", desc: "Pull's all git repo's in current directory on the master branch"},
		item{title: "Pull All - Develop", desc: "Pull's all git repo's in current directory on the develop branch"},
		item{title: "Update", desc: "Updates system and tools"},
		item{title: "Cleanup", desc: "Cleans up system"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Scripts"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}