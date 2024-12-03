package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type textModel struct {
	textInput textinput.Model
	err       error
}

var (
	placeHolderStyle = lipgloss.NewStyle().Italic(true)
	OutputValue      string
)

func textInputModel(placeHolder string) textModel {
	ti := textinput.New()
	ti.Placeholder = placeHolder
	ti.PlaceholderStyle = placeHolderStyle
	ti.Focus()
	return textModel{
		textInput: ti,
		err:       nil,
	}
}

func (m textModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m textModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			value := m.textInput.Value()
			OutputValue = value
			return m, tea.Quit
		case tea.KeyCtrlC:
			return nil, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m textModel) View() string {
	return (m.textInput.View())
}

func StartInputTextModel(placeHolder string) string {
	p := tea.NewProgram(textInputModel(placeHolder))
	p.Run()
	return OutputValue
}
