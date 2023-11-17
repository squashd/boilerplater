package ui

import "github.com/charmbracelet/lipgloss"

var (
	foregroundColour  = "#D6E5FF"
	primaryColour     = "#219ed9"
	secondaryColour   = "#70cbf4"
	mutedColour       = "#4D4D4D"
	destructiveColour = "#4D1F3B"
)

var (
	Color       = lipgloss.NewStyle().Foreground(lipgloss.Color(primaryColour))
	Choice      = lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColour)).Padding(0, 1, 0)
	Header      = lipgloss.NewStyle().Foreground(lipgloss.Color(foregroundColour)).Padding(0, 1, 0)
	Help        = lipgloss.NewStyle().Foreground(lipgloss.Color(secondaryColour)).Padding(1, 1, 0)
	Error       = lipgloss.NewStyle().Foreground(lipgloss.Color(destructiveColour)).Bold(true).Padding(0, 0, 0)
	Logo        = lipgloss.NewStyle().Foreground(lipgloss.Color(foregroundColour)).Bold(true)
	Unselected  = lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColour)).Padding(0, 1)
	Selected    = lipgloss.NewStyle().Foreground(lipgloss.Color(foregroundColour)).Padding(0, 1)
	PaddingLeft = lipgloss.NewStyle().PaddingLeft(1)
)
