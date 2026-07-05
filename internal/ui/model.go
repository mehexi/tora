package ui

import (
	"fmt"

	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mehezi/tora/internal/search"
)

var program *tea.Program

func SetProgram(p *tea.Program) {
	program = p
}

type SearchMode int

const (
	modeAll SearchMode = iota
	modeGames
	modeMovies
	modeAnime
	modeDownload
	modeSeeding
)

type ActiveWindowState int

const (
	sideBar ActiveWindowState = iota
	mainContent
	searchBar
)

func (s SearchMode) String() string {
	switch s {
	case modeAll:
		return "All"
	case modeGames:
		return "Games"
	case modeMovies:
		return "Movies"
	case modeAnime:
		return "Anime"
	case modeDownload:
		return "Download"
	case modeSeeding:
		return "Seeding"
	default:
		return "Unknown"
	}
}

func (s SearchMode) Next() SearchMode {
	switch s {
	case modeAll:
		return modeGames
	case modeGames:
		return modeMovies
	case modeMovies:
		return modeAnime
	case modeAnime:
		return modeDownload
	case modeDownload:
		return modeSeeding
	default:
		return modeAll
	}
}

func (s SearchMode) Prev() SearchMode {
	switch s {
	case modeGames:
		return modeAll
	case modeMovies:
		return modeGames
	case modeAnime:
		return modeMovies
	case modeDownload:
		return modeAnime
	case modeSeeding:
		return modeDownload
	default:
		return modeSeeding
	}
}

func (s ActiveWindowState) Next() ActiveWindowState {
	switch s {
	case sideBar:
		return mainContent
	default:
		return sideBar
	}
}
func (s ActiveWindowState) Prev() ActiveWindowState {
	switch s {
	case mainContent:
		return sideBar
	default:
		return searchBar
	}
}

type Model struct {
	appLayout  LayoutState
	appState   AppState
	normalMode NormalState
	torrent    TorrentState
}

type NormalState struct {
	ActiveWindow ActiveWindowState
}

type LayoutState struct {
	width    int
	height   int
	viewPort viewport.Model
}

type AppState struct {
	mode      SearchMode
	inputText textinput.Model
}

type TorrentState struct {
	toreentList []search.Torrent
	table       table.Model
	errMsg      string
}

func torrentsToRows(data []search.Torrent) []table.Row {
	rows := make([]table.Row, len(data))
	for i, d := range data {
		rows[i] = table.Row{
			fmt.Sprintf("%d", i+1),
			d.Name,
			d.Size,
			fmt.Sprintf("%d:%d", d.Seeders, d.Leechers),
			d.Source,
		}
	}
	return rows
}

func initTorrentTable() table.Model {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Name", Width: 50},
		{Title: "Size", Width: 10},
		{Title: "Seed:Lch", Width: 14},
		{Title: "Src", Width: 6},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(nil),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(bgColor).
		Background(purpleColor).
		Bold(false)
	t.SetStyles(s)

	return t
}

func IntialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Search torrents..."
	ti.Focus()

	vp := viewport.New()

	data, err := search.SearchEzTV()
	m := Model{
		appLayout: LayoutState{
			viewPort: vp,
		},
		appState: AppState{
			mode:      modeAll,
			inputText: ti,
		},
		torrent: TorrentState{
			toreentList: data,
			table:       initTorrentTable(),
			errMsg:      "",
		},
	}

	if err != nil {
		m.torrent.errMsg = err.Error()
	} else {
		m.torrent.table.SetRows(torrentsToRows(data))
	}

	return m
}
