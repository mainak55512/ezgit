package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type Outputs struct {
	Text_output string
}

type TuiModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	Output   Outputs
}

func InitialModel() TuiModel {
	return TuiModel{
		choices: []string{
			"Push to Remote",
			"Pull from Remote",
			"Manage Branches",
			// "Fetch from pull request",
			// "Create new Local Branch",
		},
		selected: make(map[int]struct{}),
	}
}
func (m TuiModel) Init() tea.Cmd {
	return nil
}
func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
			m.Output.Text_output = m.choices[m.cursor]
			return m, tea.Quit
		}
	}

	return m, nil
}
func (m TuiModel) View() string {
	s := "\nWhat to do?\n\n"

	for i, choice := range m.choices {

		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "✓"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
