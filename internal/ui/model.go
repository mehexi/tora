package ui

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
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

type Model struct {
	appLayout LayoutState
	appState  AppState
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

type TorentState struct {
	toreentList []Torrent
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
	}
}
