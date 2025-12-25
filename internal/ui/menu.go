package ui

import (
	"github.com/charmbracelet/bubbles/list"
)

// MenuItem represents a menu option
type MenuItem struct {
	title   string
	desc    string
	command string
}

func (i MenuItem) Title() string       { return i.title }
func (i MenuItem) Description() string { return i.desc }
func (i MenuItem) FilterValue() string { return i.title }
func (i MenuItem) Command() string     { return i.command }

// NewMenuItems returns the list of available operations
func NewMenuItems() []list.Item {
	return []list.Item{
		MenuItem{
			title:   "Show git branch",
			desc:    "Show current branch for all repos",
			command: "branch",
		},
		MenuItem{
			title:   "Pull main",
			desc:    "Pull main for all repos",
			command: "pull",
		},
	}
}

// NewMenuList creates and configures the menu list
func NewMenuList(width, height int) list.Model {
	items := NewMenuItems()
	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Scripts"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)
	return l
}
