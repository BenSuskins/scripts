package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"

	"suskins/scripts/internal/config"
	"suskins/scripts/internal/core"
)

type listKeyMap struct {
	toggleHelp       key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleTitle      key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		toggleHelp: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		togglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		toggleTitle: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
	}
}

// AppState represents the current state of the application
type AppState int

const (
	StateMenu AppState = iota
	StateExecuting
	StateResults
)

// Model is the main application model
type Model struct {
	state       AppState
	menu        list.Model
	keys        *listKeyMap
	menuTitle   string
	spinner     spinner.Model
	results     []core.OperationResult
	gitDirs     []string
	selectedCmd MenuItem
	currentIdx  int
	width       int
	height      int
}

// Messages
type gitDirsMsg struct{ dirs []string }
type singleResultMsg struct {
	result core.OperationResult
}

// NewModel creates the initial model
func NewModel(cfg *config.Config) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = SpinnerStyle

	keys := newListKeyMap()
	menu := NewMenuList(cfg, 80, 20)
	menuTitle := menu.Title

	menu.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			keys.toggleHelp,
			keys.toggleStatusBar,
			keys.togglePagination,
			keys.toggleTitle,
		}
	}

	return Model{
		state:     StateMenu,
		menu:      menu,
		keys:      keys,
		menuTitle: menuTitle,
		spinner:   s,
		width:     80,
		height:    24,
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

func executeCommand(dir string, command string) tea.Cmd {
	return func() tea.Msg {
		r := core.RunCommand(dir, command)
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
					m.selectedCmd = item
					m.state = StateExecuting
					m.currentIdx = 0
					m.results = nil

					if item.CmdType() == config.TypeSingle {
						// Single command: run once in cwd
						cwd, _ := os.Getwd()
						return m, executeCommand(cwd, item.Command())
					}

					// git_dirs: iterate over all git directories
					if len(m.gitDirs) > 0 {
						return m, executeCommand(m.gitDirs[0], item.Command())
					}
					m.state = StateResults
					return m, nil
				}
			}
			if m.state == StateResults {
				return m, tea.Quit
			}
		case "H":
			if m.state == StateMenu && !m.menu.SettingFilter() {
				m.menu.SetShowHelp(!m.menu.ShowHelp())
				return m, nil
			}
		case "S":
			if m.state == StateMenu && !m.menu.SettingFilter() {
				m.menu.SetShowStatusBar(!m.menu.ShowStatusBar())
				return m, nil
			}
		case "P":
			if m.state == StateMenu && !m.menu.SettingFilter() {
				m.menu.SetShowPagination(!m.menu.ShowPagination())
				return m, nil
			}
		case "T":
			if m.state == StateMenu && !m.menu.SettingFilter() {
				if m.menu.Title == "" {
					m.menu.Title = m.menuTitle
				} else {
					m.menu.Title = ""
				}
				return m, nil
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

		if m.selectedCmd.CmdType() == config.TypeSingle {
			// Single command done
			m.state = StateResults
			return m, nil
		}

		// git_dirs: continue to next
		m.currentIdx++
		if m.currentIdx < len(m.gitDirs) {
			return m, executeCommand(m.gitDirs[m.currentIdx], m.selectedCmd.Command())
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

	b.WriteString(TitleStyle.Render(m.selectedCmd.Title()) + "\n\n")

	// Show completed results
	for _, r := range m.results {
		icon := SuccessStyle.Render("✓")
		if !r.Success {
			icon = ErrorStyle.Render("✗")
		}
		dirName := filepath.Base(r.Directory)
		// Show output if: showOutput is true, OR command failed (always show errors)
		if (m.selectedCmd.ShowOutput() || !r.Success) && r.Message != "" {
			b.WriteString(fmt.Sprintf("%s %s: %s\n", icon, DirStyle.Render(dirName), r.Message))
		} else {
			b.WriteString(fmt.Sprintf("%s %s\n", icon, DirStyle.Render(dirName)))
		}
	}

	// For git_dirs type, show current and pending
	if m.selectedCmd.CmdType() == config.TypeGitDirs {
		// Show current repo with spinner
		if m.currentIdx < len(m.gitDirs) {
			dirName := filepath.Base(m.gitDirs[m.currentIdx])
			b.WriteString(fmt.Sprintf("%s %s\n", m.spinner.View(), dirName))
		}

		// Show pending repos
		for i := m.currentIdx + 1; i < len(m.gitDirs); i++ {
			dirName := filepath.Base(m.gitDirs[i])
			b.WriteString(PendingStyle.Render(fmt.Sprintf("  %s", dirName)) + "\n")
		}
	} else {
		// Single command: just show spinner
		b.WriteString(fmt.Sprintf("%s Running...\n", m.spinner.View()))
	}

	return b.String()
}

func (m Model) renderResults() string {
	var b strings.Builder

	b.WriteString(TitleStyle.Render(m.selectedCmd.Title()) + "\n\n")

	for _, r := range m.results {
		icon := SuccessStyle.Render("✓")
		if !r.Success {
			icon = ErrorStyle.Render("✗")
		}
		dirName := filepath.Base(r.Directory)
		// Show output if: showOutput is true, OR command failed (always show errors)
		if (m.selectedCmd.ShowOutput() || !r.Success) && r.Message != "" {
			b.WriteString(fmt.Sprintf("%s %s: %s\n", icon, DirStyle.Render(dirName), r.Message))
		} else {
			b.WriteString(fmt.Sprintf("%s %s\n", icon, DirStyle.Render(dirName)))
		}
	}

	b.WriteString("\n" + HelpStyle.Render("Press enter or q to exit"))
	return b.String()
}
