package ui

import (
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
	default:
		return modeAnime
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
	isSplashScreen bool
	mode           SearchMode
	inputText      textinput.Model
}

type TorrentState struct {
	toreentList []Torrent
	table       table.Model
}

type Torrent struct {
	Name     string
	Size     string
	Seeders  int
	Leechers int
	Magnet   string
	Category string
	Source   string
}

func newTableModel() table.Model {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Name", Width: 50},
		{Title: "Size", Width: 10},
		{Title: "Seed:Lch", Width: 14},
		{Title: "Src", Width: 6},
	}

	rows := []table.Row{
		{"1", "Obsession.2026.1080p.AMZN.WEB-DL.DDP5.1.H264.MP4-BTM", "4.96 GB", "12k:14k", "TPB"},
		{"2", "Masters of the Universe (2026) [1080p] [WEBRip]", "2.34 GB", "6714:6168", "TPB"},
		{"3", "Project Hail Mary (2026) [1080p] [WEBRip] [5.1]", "2.89 GB", "6557:1232", "TPB"},
		{"4", "House of the Dragon S03E02 1080p HEVC x265-MeGusta", "715.26 MB", "5957:3896", "TPB"},
		{"5", "Silo S03E01 1080p WEB h264-ETHEL", "3.91 GB", "5692:650", "TPB"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}

func IntialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Search torrents..."
	ti.Focus()

	vp := viewport.New()

	return Model{
		appLayout: LayoutState{
			viewPort: vp,
		},
		appState: AppState{
			isSplashScreen: true,
			mode:           modeAll,
			inputText:      ti,
		},
		torrent: TorrentState{
			table: newTableModel(),
		},
	}
}
