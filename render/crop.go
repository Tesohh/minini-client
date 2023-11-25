package render

import (
	"log/slog"
)

// Crops a matrix taking a center point (x, y) and crop size (sx, sy).
// Does not modify the original matrix, only returns a new one
//
// **NOTE** sx and sy must be even numbers or else the crop will be offset by 1
//
// https://goplay.tools/snippet/4l231LgoPwg

func (m Map) CropAndFill(x, y int, sx, sy int) Map {
	var (
		smallX = x - sx/2
		bigX   = x + sx/2
		smallY = y - sy/2
		bigY   = y + sy/2
	)

	temp := make([][]Tile, 0)
	for i := smallY; i < bigY; i++ {
		slog.Info("LETT")
		if i < 0 || i > len(m.Tiles)-1 {
			// out of bounds
			row := make([]Tile, 0)
			for ii := 0; ii > sx; ii++ {
				row = append(row, GetTile('~'))
			}

			temp = append(temp, row)
			continue
		}

		row := make([]Tile, 0)
		for j := smallX; j < bigX; j++ {
			if j < 0 || j > len(m.Tiles[i])-1 {
				row = append(row, emptyTile)
				continue
			}

			row = append(row, m.Tiles[i][j])
		}
		temp = append(temp, row)
	}

	m.Tiles = temp
	return m
}
