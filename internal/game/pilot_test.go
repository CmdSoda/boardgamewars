package game

import (
	"fmt"
	"github.com/google/uuid"
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
	assert.Nil(t, OpenDatabase())

	p := NewPilot("usa", OF5)
	assert.Nil(t, p.Save())

	CloseDatabase()
}

func TestDatabaseLoad(t *testing.T) {
	assert.Nil(t, InitGame(0))
	assert.Nil(t, OpenDatabase())
	DropPilotTable()
	assert.Nil(t, CreatePilotTable())

	p := NewPilot("usa", OF5)
	p.Name = "Gordon Link"
	p.PilotId = (PilotId)(uuid.MustParse("11977ef1-0637-4626-bbb0-598ee050270a"))
	assert.Nil(t, p.Save())
	CloseDatabase()

	assert.Nil(t, OpenDatabase())
	p2 := Pilot{}
	pid, err := uuid.Parse("11977ef1-0637-4626-bbb0-598ee050270a")
	assert.Nil(t, err)
	p2.PilotId = (PilotId)(pid)
	assert.Nil(t, p2.Load())
	assert.Equal(t, "Gordon Link", p2.Name)
	CloseDatabase()

}

func TestCreatePilotTable(t *testing.T) {
	assert.Nil(t, InitGame(0))
	assert.Nil(t, OpenDatabase())
	DropPilotTable()
	assert.Nil(t, CreatePilotTable())
	CloseDatabase()
}
