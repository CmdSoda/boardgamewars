package world

type Map struct {
	Width  int
	Height int
	Tiles  []Tile
}

//goland:noinspection ALL
// NewMap erzeugt eine Map und initialisiert alle Felder mit Default-Werten.
func NewMap(width int, height int) Map {
	m := make([]Tile, width*height, width*height)
	return Map{
		Width:  width,
		Height: height,
		Tiles:  m,
	}
}

func (m Map) GetTileAt(x int, y int) Tile {
	return m.Tiles[x-1+(m.Height-y)*m.Width]
}

func (m *Map) SetTileAt(x int, y int, t Tile) {
	m.Tiles[x-1+(m.Height-y)*m.Width] = t
}
