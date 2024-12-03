package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	SelectedBranch string
)

type AvailableBranchModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialAvailableBranchModel(branchList []string) AvailableBranchModel {
	return AvailableBranchModel{
		choices:  branchList,
		selected: make(map[int]struct{}),
	}
}
func (m AvailableBranchModel) Init() tea.Cmd {
	return nil
}
func (m AvailableBranchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			SelectedBranch = m.choices[m.cursor]
			return m, tea.Quit
		}
	}

	return m, nil
}
func (m AvailableBranchModel) View() string {
	s := "\nSelect Branch\n\n"

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

func StartAvailableBranchOptions(branchList []string) (string, error) {
	p := tea.NewProgram(InitialAvailableBranchModel(branchList))
	if _, err := p.Run(); err != nil {
		return "", err
	}
	return SelectedBranch, nil
}
