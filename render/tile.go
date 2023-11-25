package render

import (
	"github.com/charmbracelet/lipgloss"
)

type Tile struct {
	MapChar rune
	Name    string

	RenderChar rune
	FG         lipgloss.Color
	BG         lipgloss.Color

	CannotWalkOver bool
}

func (t Tile) String() string {
	return lipgloss.NewStyle().Foreground(t.FG).Background(t.BG).Render(string(t.RenderChar))
}

var unknownTile = Tile{
	MapChar:        '?',
	Name:           "unknown",
	RenderChar:     '?',
	FG:             lipgloss.Color("#f800f8"),
	CannotWalkOver: false,
}

var emptyTile = Tile{
	MapChar:        '§',
	Name:           "empty",
	RenderChar:     ' ',
	CannotWalkOver: true,
}

var tiles = []Tile{
	unknownTile,
	emptyTile,
	{
		MapChar:        '.',
		Name:           "grass",
		RenderChar:     '█',
		FG:             lipgloss.Color("#20FF20"),
		CannotWalkOver: false,
	},
	{
		MapChar:        '~',
		Name:           "water",
		RenderChar:     '~',
		FG:             lipgloss.Color("#FFFFFF"),
		BG:             lipgloss.Color("#0000FF"),
		CannotWalkOver: true,
	},
	{
		MapChar:        'S',
		Name:           "sand",
		RenderChar:     '█',
		FG:             lipgloss.Color("#C2B280"),
		CannotWalkOver: false,
	},
	{
		MapChar:        's',
		Name:           "stone",
		RenderChar:     '█',
		FG:             lipgloss.Color("#888C8D"),
		CannotWalkOver: false,
	},
	{
		MapChar:        'f',
		Name:           "flower",
		RenderChar:     '✽',
		FG:             lipgloss.Color("#888C8D"),
		BG:             lipgloss.Color("#00FF00"),
		CannotWalkOver: false,
	},
	{
		MapChar:        't',
		Name:           "tree",
		RenderChar:     '↟',
		FG:             lipgloss.Color("#FFB7C5"),
		BG:             lipgloss.Color("#00FF00"),
		CannotWalkOver: true,
	},
}

func GetTile(mapChar rune) Tile {
	for _, v := range tiles {
		if v.MapChar == mapChar {
			return v
		}
	}

	return unknownTile // unknown
}
