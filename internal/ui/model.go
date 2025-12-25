package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"suskins/scripts/internal/core"
)

// AppState represents the current state of the application
type AppState int

const (
	StateMenu AppState = iota
	StateExecuting
	StateResults
)

// Model is the main application model
type Model struct {
	state     AppState
	menu      list.Model
	results   []core.OperationResult
	gitDirs   []string
	operation string
	width     int
	height    int
}

// Messages
type gitDirsMsg struct{ dirs []string }
type resultsMsg struct{ results []core.OperationResult }

// NewModel creates the initial model
func NewModel() Model {
	return Model{
		state:  StateMenu,
		menu:   NewMenuList(80, 20),
		width:  80,
		height: 24,
	}
}

func (m Model) Init() tea.Cmd {
	return scanGitDirs
}

func scanGitDirs() tea.Msg {
	cwd, err := os.Getwd()
	if err != nil {
		return gitDirsMsg{dirs: []string{}}
	}
	dirs, _ := core.FindGitDirectories(cwd)
	return gitDirsMsg{dirs: dirs}
}

func executeOps(dirs []string, op string) tea.Cmd {
	return func() tea.Msg {
		var results []core.OperationResult
		for _, dir := range dirs {
			var r core.OperationResult
			switch op {
			case "branch":
				r = core.GetBranch(dir)
			case "pull":
				r = core.PullMain(dir)
			default:
				r = core.OperationResult{Directory: dir, Success: false, Message: "Unknown operation"}
			}
			results = append(results, r)
		}
		return resultsMsg{results: results}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.state == StateMenu {
				if item, ok := m.menu.SelectedItem().(MenuItem); ok {
					m.operation = item.Command()
					m.state = StateExecuting
					return m, executeOps(m.gitDirs, m.operation)
				}
			}
			if m.state == StateResults {
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.menu.SetSize(msg.Width, msg.Height-2)
	case gitDirsMsg:
		m.gitDirs = msg.dirs
	case resultsMsg:
		m.results = msg.results
		m.state = StateResults
	}

	if m.state == StateMenu {
		var cmd tea.Cmd
		m.menu, cmd = m.menu.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case StateMenu:
		info := HelpStyle.Render(fmt.Sprintf("Found %d git repositories", len(m.gitDirs)))
		return m.menu.View() + "\n" + info
	case StateExecuting:
		return "Executing operations..."
	case StateResults:
		return m.renderResults()
	}
	return ""
}

func (m Model) renderResults() string {
	var b strings.Builder

	opTitle := "Branch Status"
	if m.operation == "pull" {
		opTitle = "Pull Results"
	}
	b.WriteString(TitleStyle.Render(opTitle) + "\n\n")

	for _, r := range m.results {
		icon := SuccessStyle.Render("✓")
		if !r.Success {
			icon = ErrorStyle.Render("✗")
		}
		dirName := filepath.Base(r.Directory)
		b.WriteString(fmt.Sprintf("%s %s: %s\n",
			icon,
			DirStyle.Render(dirName),
			r.Message,
		))
	}

	b.WriteString("\n" + HelpStyle.Render("Press enter or q to exit"))
	return b.String()
}
