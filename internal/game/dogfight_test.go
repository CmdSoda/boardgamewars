package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CreateSetup() DogfightSetup {
	var blue = []Aircraft{
		NewAircraft("F14", "Default", WarPartyIdUSA),
		NewAircraft("F14", "Default", WarPartyIdUSA),
	}

	var red = []Aircraft{
		NewAircraft("MiG-29", "Default", WarPartyIdRussia),
		NewAircraft("MiG-29", "Default", WarPartyIdRussia),
	}

	return DogfightSetup{
		TeamBlue: AircraftIdList{blue[0].AircraftId, blue[1].AircraftId},
		TeamRed:  AircraftIdList{red[0].AircraftId, red[1].AircraftId},
	}
}

func TestDogfightSetup1(t *testing.T) {
	assert.Nil(t, InitGame())
	dfs := CreateSetup()
	fmt.Println(dfs)
}

func TestDeriveSituations(t *testing.T) {
	assert.Nil(t, InitGame())
	dfs := CreateSetup()
	fmt.Println(dfs)
}

func TestDogfight(t *testing.T) {
	assert.Nil(t, InitGame())

	fighter1 := NewAircraft("F14", "Default", WarPartyIdUK)
	fighter2 := NewAircraft("MiG-29", "Default", WarPartyIdUK)

	assert.NotNil(t, fighter1)
	assert.NotNil(t, fighter2)

	ldp1 := DogfightPositionTossup
	ldp2 := DogfightPositionTossup
	dr1, dr2 := ExecuteDogfight(fighter1.AircraftId, ldp1, fighter2.AircraftId, ldp2)
	assert.NotNil(t, dr1)
	assert.NotNil(t, dr2)
	fmt.Println(dr1)
	fmt.Println(dr2)
}

func TestMoreRounds(t *testing.T) {
	assert.Nil(t, InitGame())

	fighter1 := NewAircraft("F14", "Default", WarPartyIdUK)
	pl1 := NewPilots(fighter1.GetParameters().Seats, WarPartyIdUK, nato.OF1)
	fighter1.FillSeatsWith(pl1)
	fighter2 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
	pl2 := NewPilots(fighter2.GetParameters().Seats, WarPartyIdRussia, nato.OF1)
	fighter2.FillSeatsWith(pl2)

	assert.NotNil(t, fighter1)
	assert.NotNil(t, fighter2)

	fmt.Println(fighter1)
	fmt.Println(fighter2)

	drl1, drl2 := Sim10Rounds(fighter1.AircraftId, fighter2.AircraftId)
	assert.NotNil(t, drl1)
	assert.NotNil(t, drl2)

	fmt.Println(drl1)
	fmt.Println(drl2)
}

func TestDogfightSetup_CreateDogfight(t *testing.T) {
	assert.Nil(t, InitGame())

	b1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	r1 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)

	ds := NewDogfightSetup()
	ds.AddBlue(b1.AircraftId)
	ds.AddRed(r1.AircraftId)

	assert.Equal(t, b1.AircraftId, ds.TeamBlue[0])
	assert.Equal(t, r1.AircraftId, ds.TeamRed[0])

	d := ds.CreateDogfight()

	assert.Equal(t, 0, len(d.Groups))
	assert.Equal(t, b1.AircraftId, d.TeamBlueWaiting[0])
	assert.Equal(t, r1.AircraftId, d.TeamRedWaiting[0])

	d.DistributeAircraftsToGroups()

}
