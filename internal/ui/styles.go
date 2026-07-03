package ui

import "charm.land/lipgloss/v2"

var (
	bgColor     = lipgloss.Color("#0a0810")
	purpleColor = lipgloss.Color("#a78bfa")
	accentColor = lipgloss.Color("#b9a7e6")
	textColor   = lipgloss.Color("#e9e4f5")
	mutedColor  = lipgloss.Color("#605a6e")
	cursorColor = lipgloss.Color("#ddd8ea")

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(purpleColor)

	activeStyle = lipgloss.NewStyle().
			Foreground(purpleColor).
			Bold(true)

	inactiveStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	logoStyle = lipgloss.NewStyle().
			Foreground(cursorColor).
			Bold(true)

	logoStyleBottom = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)
)
