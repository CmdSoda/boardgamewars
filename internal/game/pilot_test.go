package game

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPilot(t *testing.T) {
	assert.Nil(t, InitGame(0))
	assert.Nil(t, RemoveAllPilots())
	p := NewPilot("usa", OF5)
	fmt.Println(p)
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
	assert.Nil(t, DropPilotTable())
	assert.Nil(t, CreatePilotTable())

	p := NewPilot("usa", OF5)
	p.Name = "Gordon Link"
	assert.Nil(t, p.Update())

	p2, errp := LoadPilot(p.PilotId)
	assert.Nil(t, errp)

	assert.Equal(t, "Gordon Link", p2.Name)

	CloseDatabase()

}

func TestCreatePilotTable(t *testing.T) {
	assert.Nil(t, InitGame(0))
	assert.Nil(t, DropPilotTable())
	assert.Nil(t, CreatePilotTable())
	CloseDatabase()
}

func TestNumberOfPilots(t *testing.T) {
	assert.Nil(t, InitGame(0))
	assert.Nil(t, RemoveAllPilots())
	NewPilot("usa", OF5)
	nr, err := NumberOfPilots()
	assert.Nil(t, err)
	assert.Equal(t, 1, nr)
}

func TestGetPilotsOfCountry(t *testing.T) {
	assert.Nil(t, InitGame(0))
	assert.Nil(t, RemoveAllPilots())
	pUsa := NewPilot("usa", OF5)
	NewPilot("uk", OF5)
	pilots, err := GetPilotsOfCountry("usa")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(pilots))
	assert.Equal(t, pUsa.Name, pilots[0].Name)
}
