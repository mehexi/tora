package ui

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func RenderSplash(m Model) string {
	ti := m.appState.inputText.View()
	bs := borderStyle.Render(ti)

	var modes []string

	for _, s := range []SearchMode{modeAll, modeGames, modeMovies, modeAnime} {
		if s == m.appState.mode {
			modes = append(modes, activeStyle.Render(s.String()))
		} else {
			modes = append(modes, inactiveStyle.Render(s.String()))
		}
	}

	modeRow := lipgloss.NewStyle().Render(strings.Join(modes, " "))

	logoTop := logoStyle.Render("▀█▀ █▀█ █▀█ █▀▀")
	logoBottom := logoStyleBottom.Render("█  █▄█ █▀▄ █▄▄")

	bottomText := logoStyleBottom.Render("A fast terminal torrent downloader")

	logo := lipgloss.JoinVertical(lipgloss.Center, logoTop, logoBottom)

	content := lipgloss.JoinVertical(lipgloss.Center, logo, "", bottomText, "", bs, modeRow)

	return lipgloss.Place(
		m.appLayout.width,
		m.appLayout.height, lipgloss.Center, lipgloss.Center, content)
}
