package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	if m.appState.isSplashScreen {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			return m.onSplashKey(msg)
		case tea.WindowSizeMsg:
			return m.onWindowSize(msg)
		}
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.onWindowSize(msg)
	}

	return m, nil
}

func (m Model) onSplashKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "tab":
		m.appState.mode = m.appState.mode.Next()
	case "shift+tab":
		m.appState.mode = m.appState.mode.Prev()
	case "enter":
		m.appState.isSplashScreen = false
	default:
		m.appState.inputText, _ = m.appState.inputText.Update(msg)
	}
	return m, nil
}

func (m Model) onWindowSize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.appLayout.width = msg.Width
	m.appLayout.height = msg.Height
	m.appState.inputText.SetWidth(m.appLayout.width - 20)
	return m, nil
}
