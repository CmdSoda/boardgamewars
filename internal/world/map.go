package world

type Map struct {
	Width int
	Height int
	Tiles []Tile
}

//goland:noinspection ALL
func NewMap(width int, height int) Map {
	m := make([]Tile, width * height, width * height)
	return Map{
		Width:  width,
		Height: height,
		Tiles:  m,
	}
}

func (m Map) GetTileAt(x int, y int) Tile {
	return m.Tiles[y * m.Width + x]
}

func (m *Map) SetTileAt(x int, y int, t Tile) {
	m.Tiles[y * m.Width + x] = t
}
