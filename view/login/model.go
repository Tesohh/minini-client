package login

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"github.com/Tesohh/minini-client/connection"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	errorStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color(lipgloss.Color("#FF0000")))
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy()
	blurredButton = blurredStyle.Copy() //fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type Model struct {
	focusIndex int
	inputs     []textinput.Model

	// login or signup
	mode           string
	wrongSomething bool

	c *connection.Client
	s *connection.ServerConn
}

func InitialModel(c *connection.Client, s *connection.ServerConn, mode string) Model {
	m := Model{
		inputs: make([]textinput.Model, 2),
		c:      c,
		s:      s,
		mode:   mode,
	}

	m.inputs[0] = textinput.New()
	m.inputs[0].Cursor.Style = cursorStyle
	m.inputs[0].CharLimit = 12
	m.inputs[0].Placeholder = "Username"
	m.inputs[0].Focus()
	m.inputs[0].PromptStyle = focusedStyle
	m.inputs[0].TextStyle = focusedStyle

	m.inputs[1] = textinput.New()
	m.inputs[1].Cursor.Style = cursorStyle
	m.inputs[1].CharLimit = 20
	m.inputs[1].Placeholder = "Password"
	m.inputs[1].EchoMode = textinput.EchoPassword
	m.inputs[1].EchoCharacter = '•'

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
