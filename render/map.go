package render

import (
	"bytes"
	"os"
	"path"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

var biomeColors = map[string]lipgloss.Color{
	"forest": GetTile('.').BG,
}

type Map struct {
	Name  string `yaml:"name,omitempty"`
	Biome string `yaml:"biome,omitempty"`

	Tiles [][]Tile
}

func (m Map) String() string {
	s := ""
	for _, row := range m.Tiles {
		for _, col := range row {
			if col.BG == "" {
				col.BG = biomeColors[m.Biome]
			}

			s += col.String()
		}
		s += "\n"
	}
	return strings.TrimSuffix(s, "\n")
}

func MapFromFile(filename string) (*Map, error) {
	file, err := os.Open(path.Clean(filename))
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 2048)
	n, err := file.Read(buf)
	if err != nil {
		return nil, err
	}

	buf = buf[:n]

	split := bytes.Split(buf, []byte("---"))

	m := Map{}
	m.Tiles = make([][]Tile, 0)
	metaData := split[0]
	err = yaml.Unmarshal(metaData, &m)
	if err != nil {
		return nil, err
	}

	mapData := bytes.Trim(split[1], "\n")
	lines := bytes.Split(mapData, []byte("\n"))

	for i, line := range lines {
		runebuf := bytes.Runes(line)
		m.Tiles = append(m.Tiles, make([]Tile, 0))
		for _, r := range runebuf {
			m.Tiles[i] = append(m.Tiles[i], GetTile(r))
		}
	}

	return &m, nil
}
