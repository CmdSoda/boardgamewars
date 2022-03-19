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

func TestDatabaseSave(t *testing.T) {
	assert.Nil(t, InitGame(0))

	p := NewPilot("usa", OF5)
	fmt.Println(p.PilotId)

	CloseDatabase()
}

func TestDatabaseLoad(t *testing.T) {
	assert.Nil(t, InitGame(0))
	DropPilotTable()
	assert.Nil(t, CreatePilotTable())

	p := NewPilot("usa", OF5)
	p.Name = "Gordon Link"
	assert.Nil(t, p.Update())

	p2 := Pilot{}
	p2.PilotId = p.PilotId
	assert.Nil(t, p2.Load())
	assert.Equal(t, "Gordon Link", p2.Name)

	CloseDatabase()

}

func TestCreatePilotTable(t *testing.T) {
	assert.Nil(t, InitGame(0))
	DropPilotTable()
	assert.Nil(t, CreatePilotTable())
	CloseDatabase()
}
