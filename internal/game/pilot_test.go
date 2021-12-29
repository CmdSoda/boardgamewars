package game

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPilot(t *testing.T) {
	assert.Nil(t, InitGame(0))
	p := NewPilot("usa", OF5)
	fmt.Println(p)
	p2, exist := Globals.AllPilots[p.PilotId]
	assert.True(t, exist)
	assert.Equal(t, p2, p)
}

func TestNewPilots(t *testing.T) {
	assert.Nil(t, InitGame(0))
	pl := NewPilots(3, "usa", OF1)
	fmt.Println(pl)
}
