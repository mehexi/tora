package ui

import "charm.land/lipgloss/v2"

var (
	bgColor     = lipgloss.Color("#292928")
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
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			PaddingLeft(1).
			BorderLeftForeground(purpleColor).
			Bold(true)

	inactiveStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			BorderLeftForeground(bgColor).
			PaddingLeft(1).
			Foreground(mutedColor)

	logoStyle = lipgloss.NewStyle().
			Foreground(cursorColor).
			Bold(true)

	logoStyleBottom = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)

	mutedActive = lipgloss.NewStyle().
			Foreground(purpleColor).
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			PaddingLeft(1).
			BorderLeftForeground(mutedColor).
			Bold(true)

	footerIndicator = lipgloss.NewStyle().
			Background(purpleColor).
			Foreground(bgColor).
			Bold(true).
			Padding(0, 1)

	footerMuted = lipgloss.NewStyle().
			Foreground(mutedColor)
)
