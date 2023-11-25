package game

import (
	"github.com/Tesohh/minini-client/connection"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	frameStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Width(50).Height(25)
)

type Model struct {
	c *connection.Client
	s *connection.ServerConn
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel(c *connection.Client, s *connection.ServerConn) Model {
	return Model{
		c: c,
		s: s,
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	s := ""
	s += m.c.Username
	return frameStyle.Render(s)
}

func (m Model) Keys() help.KeyMap {
	return keys
}
