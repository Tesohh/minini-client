package rp

import (
	tea "github.com/charmbracelet/bubbletea"
)

var Global Repo

type Repo struct {
	TeaProg *tea.Program
}
