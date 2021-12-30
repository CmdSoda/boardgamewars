package game

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEventList(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	AppendEvent(NewEvent(uuid.New(), SourceTypeAircraft, "nice"))
	assert.Equal(t, 1, len(Globals.EventList))
	assert.Equal(t, StepTime(0), Globals.World.CurrentStep)
	assert.Equal(t, "nice", Globals.EventList[0].Message)
	Globals.World.Step(4)
	AppendEvent(NewEvent(uuid.New(), SourceTypeAircraft, "cool"))
	assert.Equal(t, 2, len(Globals.EventList))
	assert.Equal(t, StepTime(4), Globals.World.CurrentStep)
	assert.Equal(t, "nice", Globals.EventList[0].Message)
	assert.Equal(t, StepTime(0), Globals.EventList[0].StepTime)
	assert.Equal(t, "cool", Globals.EventList[1].Message)
	assert.Equal(t, StepTime(4), Globals.EventList[1].StepTime)
	fmt.Println(Globals.EventList.String())
}

func TestAppendEvent(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac1 := NewAircraft("F14", "Default", "usa")
	ac1.FillSeatsWithNewPilots(OF1)
	assert.Equal(t, 1, len(Globals.EventList))
	e := Globals.EventList[0]
	assert.Equal(t, "created", e.Message)
	assert.Equal(t, ac1.AircraftId, AircraftId(e.Source))
	assert.Equal(t, SourceTypeAircraft, e.SourceType)
	assert.Equal(t, StepTime(0), e.StepTime)
}
