package hud

import (
	"log/slog"
	"strings"

	"github.com/Tesohh/minini-client/connection"
	"github.com/Tesohh/minini-client/view/hud/game"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type teaHudModel interface {
	tea.Model
	Keys() help.KeyMap
}

type Model struct {
	c *connection.Client
	s *connection.ServerConn

	models          map[string]teaHudModel
	focusableModels []string
	focusIndex      int
	// keystrokes are forwarded to the current focus

	help help.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel(c *connection.Client, s *connection.ServerConn) Model {
	return Model{
		c: c,
		s: s,
		models: map[string]teaHudModel{
			"game": game.InitialModel(c, s),
		},
		focusableModels: []string{"game"},

		help: help.New(),
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	}
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder

	currentModel := m.models[m.focusableModels[m.focusIndex]]
	slog.Info("model", "model", currentModel.Keys())

	b.WriteString(m.models["game"].View())

	b.WriteByte('\n')
	b.WriteString(m.help.View(currentModel.Keys()))

	return b.String()
}
