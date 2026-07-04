package ui

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func RenderNormal(m Model) string {
	body := renderBody(m)
	return body
}

func renderLogo(m Model) string {
	logoTop := logoStyle.Render("▀█▀ █▀█ █▀█ █▀▀")
	logoBottom := logoStyleBottom.Render("█  █▄█ █▀▄ █▄▄")

	logo := lipgloss.JoinVertical(lipgloss.Center, logoTop, logoBottom)
	content := lipgloss.NewStyle().
		BorderBottom(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(mutedColor).
		Width(m.appLayout.width).
		Render(logo)

	return content
}

func renderSearchBar(m Model, width int) string {
	inputStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(mutedColor).
		Width(width)

	searchBox := strings.TrimRight(m.appState.inputText.View(), "\n")
	searchInput := inputStyle.Render(searchBox)

	return searchInput
}

func renderBody(m Model) string {
	logo := renderLogo(m)
	sideBar := renderSideBar(m)
	footer := renderedFotter(m)

	w, h := RemainingVertical(m.appLayout.width, m.appLayout.height, logo, " ", footer)
	w, h = RemainingHorizontal(w, h, sideBar)

	searchBar := renderSearchBar(m, w)

	w2, h2 := RemainingVertical(w, h, searchBar)

	searchItems := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(mutedColor).
		Width(w2).
		Height(h2).
		Render("text")

	mainContent := lipgloss.JoinVertical(lipgloss.Left, searchBar, searchItems)
	main := lipgloss.JoinHorizontal(lipgloss.Left, sideBar, mainContent)

	return lipgloss.JoinVertical(lipgloss.Left, logo, " ", main, footer)
}

func renderSideBar(m Model) string {

	var modes []string

	for _, s := range []SearchMode{modeAll, modeGames, modeMovies, modeAnime} {
		if s == m.appState.mode {
			modes = append(modes, activeStyle.Render(s.String()))
		} else {
			modes = append(modes, inactiveStyle.Render(s.String()))
		}
	}

	if m.normalMode.ActiveWindow == sideBar {
		return lipgloss.NewStyle().Width(m.appLayout.width / 9).Render(strings.Join(modes, "\n"))
	}

	return lipgloss.NewStyle().Width(m.appLayout.width / 9).Render(strings.Join(modes, "\n"))

}

func renderedFotter(m Model) string {
	return lipgloss.JoinVertical(lipgloss.Center, "mode")
}

// RemainingVertical returns the leftover width/height after subtracting
// the rendered size of the given elements (stacked vertically).
func RemainingVertical(totalWidth, totalHeight int, elements ...string) (width, height int) {
	height = totalHeight
	for _, el := range elements {
		height -= lipgloss.Height(el)
	}
	if height < 0 {
		height = 0
	}
	return totalWidth, height
}

// RemainingHorizontal returns the leftover width/height after subtracting
// the rendered size of the given elements (placed side by side).
func RemainingHorizontal(totalWidth, totalHeight int, elements ...string) (width, height int) {
	width = totalWidth
	for _, el := range elements {
		width -= lipgloss.Width(el)
	}
	if width < 0 {
		width = 0
	}
	return width, totalHeight
}
