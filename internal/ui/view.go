package ui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	v := tea.NewView(RenderNormal(m))
	v.AltScreen = true
	v.BackgroundColor = lipgloss.Black
	return v
}
