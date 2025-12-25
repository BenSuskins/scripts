package ui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	DirStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	HelpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)
