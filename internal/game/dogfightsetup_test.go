package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CreateSetup() DogfightSetup {
	var blue = []*Aircraft{
		NewAircraftManned("F14", "Default", countrycodes.USA, nato.OF1),
		NewAircraftManned("F14", "Default", countrycodes.USA, nato.OF1),
	}

	var red = []*Aircraft{
		NewAircraftManned("MiG-29", "Default", countrycodes.Russia, nato.OF1),
		NewAircraftManned("MiG-29", "Default", countrycodes.Russia, nato.OF1),
	}

	return DogfightSetup{
		TeamBlue: []AircraftId{ blue[0].AircraftId, blue[1].AircraftId },
		TeamRed:  []AircraftId{ red[0].AircraftId, red[1].AircraftId },
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