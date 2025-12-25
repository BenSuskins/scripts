package ui

import (
	"github.com/charmbracelet/bubbles/list"

	"suskins/scripts/internal/config"
)

// MenuItem represents a menu option
type MenuItem struct {
	title      string
	desc       string
	command    string
	cmdType    config.CommandType
	showOutput bool
}

func (i MenuItem) Title() string               { return i.title }
func (i MenuItem) Description() string         { return i.desc }
func (i MenuItem) FilterValue() string         { return i.title }
func (i MenuItem) Command() string             { return i.command }
func (i MenuItem) CmdType() config.CommandType { return i.cmdType }
func (i MenuItem) ShowOutput() bool            { return i.showOutput }

// NewMenuItems returns the list of available operations from config
func NewMenuItems(cfg *config.Config) []list.Item {
	items := make([]list.Item, len(cfg.Commands))
	for i, cmd := range cfg.Commands {
		items[i] = MenuItem{
			title:      cmd.Name,
			desc:       cmd.Description,
			command:    cmd.Command,
			cmdType:    cmd.Type,
			showOutput: cmd.ShowOutput,
		}
	}
	return items
}

// NewMenuList creates and configures the menu list
func NewMenuList(cfg *config.Config, width, height int) list.Model {
	items := NewMenuItems(cfg)
	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Scripts"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
	return l
}
