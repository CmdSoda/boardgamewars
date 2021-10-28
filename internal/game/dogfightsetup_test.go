package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDogfightSetup1(t *testing.T) {
	assert.Nil(t, InitGame())

	var blue = []Aircraft{
		NewAircraftManned("F14", "Default", countrycodes.USA, nato.OF1),
		NewAircraftManned("F14", "Default", countrycodes.USA, nato.OF1),
	}

	var red = []Aircraft{
		NewAircraftManned("MiG-29", "Default", countrycodes.Russia, nato.OF1),
		NewAircraftManned("MiG-29", "Default", countrycodes.Russia, nato.OF1),
	}

	dfs := NewDogfightSetup(blue, red)

	fmt.Println(dfs)
}
