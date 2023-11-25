package game

import (
	"github.com/Tesohh/minini-client/connection"
	"github.com/Tesohh/minini-client/render"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	frameStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Width(50).Height(25).PaddingLeft(1)
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Up):
			m.c.State.Y -= 1
		case key.Matches(msg, keys.Down):
			m.c.State.Y += 1
		case key.Matches(msg, keys.Left):
			m.c.State.X -= 1
		case key.Matches(msg, keys.Right):
			m.c.State.X += 1
		}
		return m, nil
	}
	return m, nil
}

func (m Model) View() string {
	s := ""
	s += m.c.Username
	s += "\n"
	gm, err := render.MapFromFile("assets/maps/world1.txt")
	if err == nil {
		gm := gm.CropAndFill(m.c.State.X, m.c.State.Y, 48, 24)
		s += gm.String()
	}
	return frameStyle.Render(s)
}

func (m Model) Keys() help.KeyMap {
	return keys
}
