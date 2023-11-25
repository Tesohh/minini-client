package login

import (
	"github.com/Tesohh/minini-client/connection"
	"github.com/Tesohh/minini-client/message"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *Model) SetSomethingWrong() {
	m.wrongSomething = true
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			caught := false
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				caught = true
				m.s.Send(message.Msg{
					Action: m.mode,
					Data:   map[string]any{"username": m.inputs[0].Value(), "password": m.inputs[1].Value()},
				})
				m.c.Username = m.inputs[0].Value()
				// return m, tea.Quit
			}

			if s == "enter" && m.focusIndex == len(m.inputs)+1 {
				caught = true
				if m.mode == "login" {
					m.mode = "signup"
				} else if m.mode == "signup" {
					m.mode = "login"
				}
				// return m, tea
			}

			if !caught {
				// Cycle indexes
				if s == "up" || s == "shift+tab" {
					m.focusIndex--
				} else {
					m.focusIndex++
				}

				if m.focusIndex > len(m.inputs)+1 {
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = len(m.inputs) + 1
				}

				cmds := make([]tea.Cmd, len(m.inputs))
				for i := 0; i <= len(m.inputs)-1; i++ {
					if i == m.focusIndex {
						// Set focused state
						cmds[i] = m.inputs[i].Focus()
						m.inputs[i].PromptStyle = focusedStyle
						m.inputs[i].TextStyle = focusedStyle
						continue
					}
					// Remove focused state
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = noStyle
					m.inputs[i].TextStyle = noStyle
				}

				return m, tea.Batch(cmds...)
			}
		}

	case connection.OkMsg:
		if bool(msg) {
			return m, tea.Quit
		} else {
			m.SetSomethingWrong()
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}
