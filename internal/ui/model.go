package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
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
	state      AppState
	menu       list.Model
	spinner    spinner.Model
	results    []core.OperationResult
	gitDirs    []string
	operation  string
	currentIdx int
	width      int
	height     int
}

// Messages
type gitDirsMsg struct{ dirs []string }
type singleResultMsg struct {
	result core.OperationResult
}

// NewModel creates the initial model
func NewModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = SpinnerStyle

	return Model{
		state:   StateMenu,
		menu:    NewMenuList(80, 20),
		spinner: s,
		width:   80,
		height:  24,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(scanGitDirs, m.spinner.Tick)
}

func scanGitDirs() tea.Msg {
	cwd, err := os.Getwd()
	if err != nil {
		return gitDirsMsg{dirs: []string{}}
	}
	dirs, _ := core.FindGitDirectories(cwd)
	return gitDirsMsg{dirs: dirs}
}

func executeNextOp(dir string, op string) tea.Cmd {
	return func() tea.Msg {
		var r core.OperationResult
		switch op {
		case "branch":
			r = core.GetBranch(dir)
		case "pull":
			r = core.PullMain(dir)
		default:
			r = core.OperationResult{Directory: dir, Success: false, Message: "Unknown operation"}
		}
		return singleResultMsg{result: r}
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
					m.currentIdx = 0
					m.results = nil
					if len(m.gitDirs) > 0 {
						return m, executeNextOp(m.gitDirs[0], m.operation)
					}
					m.state = StateResults
					return m, nil
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
	case singleResultMsg:
		m.results = append(m.results, msg.result)
		m.currentIdx++
		if m.currentIdx < len(m.gitDirs) {
			return m, executeNextOp(m.gitDirs[m.currentIdx], m.operation)
		}
		m.state = StateResults
		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
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
		return m.renderExecuting()
	case StateResults:
		return m.renderResults()
	}
	return ""
}

func (m Model) renderExecuting() string {
	var b strings.Builder

	opTitle := "Branch Status"
	if m.operation == "pull" {
		opTitle = "Pull Results"
	}
	b.WriteString(TitleStyle.Render(opTitle) + "\n\n")

	// Show completed results
	for _, r := range m.results {
		icon := SuccessStyle.Render("✓")
		if !r.Success {
			icon = ErrorStyle.Render("✗")
		}
		dirName := filepath.Base(r.Directory)
		if r.Message != "" {
			b.WriteString(fmt.Sprintf("%s %s: %s\n", icon, DirStyle.Render(dirName), r.Message))
		} else {
			b.WriteString(fmt.Sprintf("%s %s\n", icon, DirStyle.Render(dirName)))
		}
	}

	// Show current repo with spinner
	if m.currentIdx < len(m.gitDirs) {
		dirName := filepath.Base(m.gitDirs[m.currentIdx])
		b.WriteString(fmt.Sprintf("%s %s\n",
			m.spinner.View(),
			dirName,
		))
	}

	// Show pending repos
	for i := m.currentIdx + 1; i < len(m.gitDirs); i++ {
		dirName := filepath.Base(m.gitDirs[i])
		b.WriteString(PendingStyle.Render(fmt.Sprintf("  %s", dirName)) + "\n")
	}

	return b.String()
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
		if r.Message != "" {
			b.WriteString(fmt.Sprintf("%s %s: %s\n", icon, DirStyle.Render(dirName), r.Message))
		} else {
			b.WriteString(fmt.Sprintf("%s %s\n", icon, DirStyle.Render(dirName)))
		}
	}

	b.WriteString("\n" + HelpStyle.Render("Press enter or q to exit"))
	return b.String()
}
