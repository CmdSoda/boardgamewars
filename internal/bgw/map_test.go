package bgw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMap(t *testing.T) {
	m := NewMap(3, 4)
	assert.Equal(t, 12, len(m.Tiles))
	assert.Equal(t, 12, cap(m.Tiles))
	assert.Equal(t, 3, m.Width)
	assert.Equal(t, 4, m.Height)
}

func TestGetTileAt(t *testing.T) {
	m := NewMap(3, 3)
	m.SetTileAt(Position{1, 1}, Tile{
		HillType:   HillPlain,
		GroundType: GroundMud,
	})
	assert.Equal(t, Tile{
		HillType:   HillPlain,
		GroundType: GroundMud,
	}, m.GetTileAt(Position{1, 1}))
	assert.Equal(t, Tile{
		HillType:   HillPlain,
		GroundType: GroundMud,
	}, m.Tiles[6])
	assert.Equal(t, Tile{
		HillType:   HillPlain,
		GroundType: GroundStreet,
	}, m.GetTileAt(Position{0, 0}))
}

func TestSearch(t *testing.T) {
	m := NewMap(3, 3)
	m.SearchDeep(Position{
		Column: 1,
		Row:    1,
	}, Position{
		Column: 3,
		Row:    3,
	}, nil)

}

func Test_searchWide(t *testing.T) {
	m := NewMap(10, 10)
	start := Position{
		Column: 1,
		Row:    1,
	}
	end := Position{
		Column: 10,
		Row:    10,
	}
	currentpath := make(PositionList, 0, 0)
	currentpath = append(currentpath, start)
	results := make(ResultList, 0, 0)
	m.recursiveDeep(currentpath, end, &results)

	for _, r := range results {
		assert.Equal(t, end, r.Path.Last())
	}
}