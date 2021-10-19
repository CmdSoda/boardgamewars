package bgw

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

func (m Map) GetTileAt(p Position) Tile {
	return m.Tiles[p.Column-1+(m.Height-p.Row)*m.Width]
}

func (m *Map) SetTileAt(p Position, t Tile) {
	m.Tiles[p.Column-1+(m.Height-p.Row)*m.Width] = t
}

func (m *Map) InsideMap(p Position) bool {
	return p.Column >= 1 && p.Column <= m.Width && p.Row >= 1 && p.Row <= m.Height
}

type SearchHandler func(t Tile, object *interface{}) float32

type SearchParameter struct {
	Object interface{}
	Handlers []SearchHandler
}

func (m Map) recursiveDeep(path PositionList, end Position, results *ResultList) {
	if path.Last().Equal(end) {
		co := make(PositionList, len(path))
		copy(co, path)
		*results = append(*results, Result{Path: co})
		return
	}

	// Terminierung, falls zu lang
	if len(path) > m.Width + m.Height {
		return
	}

	for d := NW; d <= W; d++ {
		adj := path.Last().GetAdjacent(d)
		if adj != nil && m.InsideMap(*adj) && !path.Contains(*adj){
			path2 := append(path, *adj)
			m.recursiveDeep(path2, end, results)
		}
	}
}

func (m Map) recursiveWide(path PositionList, end Position, results *ResultList) {

}

func (m Map) SearchDeep(start Position, end Position, sp *SearchParameter) PositionList {
	currentPath := make(PositionList, 0, 0)
	currentPath = append(currentPath, start)
	results := make(ResultList, 0, 0)
	m.recursiveDeep(currentPath, end, &results)
	return currentPath
}

func (m Map) SearchWide(start Position, end Position, sp *SearchParameter) PositionList {
	currentPath := make(PositionList, 0, 0)
	results := make(ResultList, 0, 0)
	m.recursiveWide(currentPath, end, &results)
	return currentPath
}
