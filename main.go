package main

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/mainak55512/ezgit/handler"
	"github.com/mainak55512/ezgit/tui"
	"os"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	res, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	} else {
		if m, ok := res.(tui.TuiModel); ok {
			// m.Output.RunCommands()
			handler.Handler(m.Output)
		}
	}
}
