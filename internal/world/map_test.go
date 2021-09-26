package world

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMap(t *testing.T) {
	m := NewMap(3, 3)
	assert.Equal(t, 9, len(m.Tiles))
	assert.Equal(t, 9, cap(m.Tiles))
}

func TestGetTileAt(t *testing.T) {
	m := NewMap(3, 3)
	m.SetTileAt(1, 1, Tile{
		HillType:   HillPlain,
		GroundType: GroundMud,
	})
	assert.Equal(t, Tile{
		HillType:   HillPlain,
		GroundType: GroundMud,
	}, m.GetTileAt(1, 1))
	assert.Equal(t, Tile{
		HillType:   HillPlain,
		GroundType: GroundMud,
	}, m.Tiles[4])
	assert.Equal(t, Tile{
		HillType:   HillPlain,
		GroundType: GroundStreet,
	}, m.GetTileAt(0, 0))
}

