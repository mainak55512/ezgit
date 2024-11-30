package main

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"mainak55512/ezgit/config"
	"mainak55512/ezgit/tui"
	"os"
)

func main() {
	if err := config.EZInit(); err != nil {
		panic(err)
	}
	p := tea.NewProgram(tui.InitialModel())
	res, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	} else {
		if m, ok := res.(tui.TuiModel); ok {
			m.Output.RunCommands()
		}
	}
}
