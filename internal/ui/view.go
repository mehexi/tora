package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	v := tea.NewView(RenderNormal(m))
	v.AltScreen = true
	v.BackgroundColor = bgColor

	if m.appState.isSplashScreen {
		v.SetContent(RenderSplash(m))
	}

	return v
}
