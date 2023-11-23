package login

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(m.mode)
	if m.wrongSomething {
		b.WriteString(errorStyle.Render(" Wrong username or password"))
	}
	b.WriteString("\n")

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := blurredButton.Render("[ Submit ]")
	if m.focusIndex == len(m.inputs) {
		button = focusedButton.Render("[ Submit ]")
	}
	fmt.Fprintf(&b, "\n\n%s", button)

	button = blurredButton.Render("[ Switch ]")
	if m.focusIndex == len(m.inputs)+1 {
		button = focusedButton.Render("[ Switch ]")
	}
	fmt.Fprintf(&b, " %s\n\n", button)

	return b.String()
}
