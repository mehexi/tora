package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mehezi/tora/internal/ui"
)

func main() {
	m := ui.IntialModel()
	p := tea.NewProgram(m)
	ui.SetProgram(p)

	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
