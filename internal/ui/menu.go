package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"

	"suskins/scripts/internal/config"
)

// MenuItem represents a menu option
type MenuItem struct {
	title      string
	desc       string
	command    string
	cmdType    config.CommandType
	category   string
	showOutput bool
}

func (i MenuItem) Title() string               { return i.title }
func (i MenuItem) Description() string         { return i.desc }
func (i MenuItem) FilterValue() string         { return i.title }
func (i MenuItem) Command() string             { return i.command }
func (i MenuItem) CmdType() config.CommandType { return i.cmdType }
func (i MenuItem) Category() string            { return i.category }
func (i MenuItem) ShowOutput() bool            { return i.showOutput }

// NewMenuItems returns the list of available operations from config
func NewMenuItems(cfg *config.Config) []list.Item {
	items := make([]list.Item, len(cfg.Commands))
	for i, cmd := range cfg.Commands {
		items[i] = MenuItem{
			title:      fmt.Sprintf("[%s] %s", cmd.Category, cmd.Name),
			desc:       cmd.Command,
			command:    cmd.Command,
			cmdType:    cmd.Type,
			category:   cmd.Category,
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
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)
	return l
}
