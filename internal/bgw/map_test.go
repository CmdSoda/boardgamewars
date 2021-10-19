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
	results := make(SolutionList, 0, 0)
	m.recursiveDeep(currentpath, end, &results)

	for _, r := range results {
		assert.Equal(t, end, r.Path.LastElement())
	}
}

func Test_arrayCopy(t *testing.T) {
	a1 := make([]int, 0) // { }
	a1 = append(a1, 1) // { 1 }
	a1 = append(a1, 2) // { 1, 2 }
	a1 = append(a1, 3) // { 1, 2, 3 }

	a2 := make([]int, len(a1)) // { 0, 0, 0 }
	copy(a2, a1) // { 1, 2, 3 }

	a2 = append(a2, 4)
	a2[0] = 5

	if len(a1) != 3 {
		t.Log("len not 3")
		t.FailNow()
	}
	if a1[0] != 1 {
		t.Log("value not 1")
		t.FailNow()
	}

	if len(a2) != 4 {
		t.Log("len not 4")
		t.FailNow()
	}
	if a2[0] != 5 {
		t.Log("value not 5")
		t.FailNow()
	}
}

func Test_evolveSolution(t *testing.T) {
	m := NewMap(5, 5)
	m.SearchWide(Position{
		Column: 1,
		Row:    1,
	}, Position{
		Column: 3,
		Row:    3,
	}, nil)
}