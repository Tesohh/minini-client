package game

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
}

func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{}
}
func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Left, k.Down, k.Up, k.Right},
	}
}

var keys = keymap{
	Up: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "down"),
	),
	Left: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "left"),
	),
	Right: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "right"),
	),
}
