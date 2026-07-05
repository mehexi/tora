package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/mehezi/tora/internal/search"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		return m.onNormalKey(msg)
	case tea.WindowSizeMsg:
		return m.onWindowSize(msg)
	}

	return m, nil
}

func (m Model) onNormalKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.normalMode.ActiveWindow {
	case sideBar:
		switch msg.String() {
		case "j", "down":
			m.appState.mode = m.appState.mode.Next()
		case "k", "up":
			m.appState.mode = m.appState.mode.Prev()
		}
	case searchBar:
		if msg.String() == "esc" {
			m.normalMode.ActiveWindow = mainContent
		} else if msg.String() == "enter" {
			q := m.appState.inputText.Value()
			data, err := search.SearchYTS(q)
			if err != nil {
				m.torrent.errMsg = err.Error()
			} else {
				m.torrent.toreentList = data
				m.torrent.table.SetRows(torrentsToRows(data))
				m.torrent.errMsg = ""
			}
			m.normalMode.ActiveWindow = mainContent
		} else {
			m.appState.inputText, _ = m.appState.inputText.Update(msg)
		}
	case mainContent:
		m.torrent.table, cmd = m.torrent.table.Update(msg)
	}

	switch msg.String() {
	case "tab":
		m.normalMode.ActiveWindow = m.normalMode.ActiveWindow.Next()
	case "/":
		m.normalMode.ActiveWindow = searchBar
	}

	return m, cmd
}

func (m Model) onWindowSize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.appLayout.width = msg.Width
	m.appLayout.height = msg.Height
	m.appState.inputText.SetWidth(m.appLayout.width - 20)
	return m, nil
}
