package game

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CreateSetup() DogfightSetup {
	var blue = []*Aircraft{
		NewAircraft("F14", "Default", WarPartyIdUSA),
		NewAircraft("F14", "Default", WarPartyIdUSA),
	}

	var red = []*Aircraft{
		NewAircraft("MiG-29", "Default", WarPartyIdRussia),
		NewAircraft("MiG-29", "Default", WarPartyIdRussia),
	}

	return DogfightSetup{
		TeamBlue: []AircraftId{blue[0].AircraftId, blue[1].AircraftId},
		TeamRed:  []AircraftId{red[0].AircraftId, red[1].AircraftId},
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
