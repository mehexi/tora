package ui

import (
	"strings"

	"charm.land/bubbles/v2/table"
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

	if m.normalMode.ActiveWindow == searchBar {
		inputStyle = inputStyle.BorderForeground(purpleColor)
	}

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

	const borderOverhead = 3
	tableWidth := w2 - borderOverhead
	tableHeight := h2 - borderOverhead

	torrentList := renderTorrentlist(m, tableWidth, tableHeight)

	searchItemsStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(mutedColor).
		Width(w2).
		Height(h2)

	if m.normalMode.ActiveWindow == mainContent {
		searchItemsStyle = searchItemsStyle.BorderForeground(purpleColor)
	}

	searchItems := searchItemsStyle.Render(torrentList)

	mainContent := lipgloss.JoinVertical(lipgloss.Left, searchBar, searchItems)
	main := lipgloss.JoinHorizontal(lipgloss.Left, sideBar, mainContent)

	return lipgloss.JoinVertical(lipgloss.Left, logo, " ", main, footer)
}

func renderSideBar(m Model) string {
	var modes []string
	for _, s := range []SearchMode{modeAll, modeGames, modeMovies, modeAnime} {
		if s == m.appState.mode {
			if m.normalMode.ActiveWindow == sideBar {
				modes = append(modes, activeStyle.Render(s.String()))
			} else {
				modes = append(modes, mutedActive.Render(s.String()))
			}
		} else {
			modes = append(modes, inactiveStyle.Render(s.String()))
		}
	}

	return lipgloss.NewStyle().
		Width(m.appLayout.width / 9).
		Render(strings.Join(modes, "\n"))

}

func renderedFotter(m Model) string {
	var indicator string
	var help string
	switch m.normalMode.ActiveWindow {
	case sideBar:
		indicator = "SIDEBAR"
		help = "  j/k navigate • / search • ? help"
	case searchBar:
		indicator = "SEARCH"
		help = "  type query • enter search • esc back"
	case mainContent:
		indicator = "RESULTS"
		help = "  j/k scroll • enter open • / search • ? help"
	}

	label := lipgloss.JoinHorizontal(lipgloss.Center,
		footerIndicator.Render(indicator),
		footerMuted.Render(help),
	)

	return lipgloss.NewStyle().
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(mutedColor).
		Width(m.appLayout.width).
		Render(label)
}

func renderTorrentlist(m Model, w int, h int) string {
	columns := buildColumns(w)
	m.torrent.table.SetColumns(columns)
	m.torrent.table.SetWidth(w)
	m.torrent.table.SetHeight(h)
	return m.torrent.table.View()
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

func buildColumns(totalWidth int) []table.Column {
	const (
		numW  = 4
		sizeW = 10
		seedW = 14
		srcW  = 6

		numColumns    = 5
		paddingPerCol = 2 // typical: 1 space left + 1 space right, per column style
	)

	fixedTotal := numW + sizeW + seedW + srcW
	overhead := numColumns * paddingPerCol

	nameW := totalWidth - fixedTotal - overhead
	if nameW < 10 {
		nameW = 10
	}

	return []table.Column{
		{Title: "#", Width: numW},
		{Title: "Name", Width: nameW},
		{Title: "Size", Width: sizeW},
		{Title: "Seed:Lch", Width: seedW},
		{Title: "Src", Width: srcW},
	}
}
