package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	SelectedBranchOption string // Output from the branch operations tui
)

type BranchModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialBranchModel() BranchModel {
	return BranchModel{
		choices: []string{
			"Switch Branch",
			"Create & Switch Branch",
			"Delete Branch",
		},
		selected: make(map[int]struct{}),
	}
}
func (m BranchModel) Init() tea.Cmd {
	return nil
}
func (m BranchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			SelectedBranchOption = m.choices[m.cursor]
			return m, tea.Quit
		}
	}

	return m, nil
}
func (m BranchModel) View() string {
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

func StartBranchModel() (string, error) {
	p := tea.NewProgram(InitialBranchModel())
	if _, err := p.Run(); err != nil {
		return "", err
	}
	return SelectedBranchOption, nil
}
