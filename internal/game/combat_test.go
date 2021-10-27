package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDogfight(t *testing.T) {
	InitGame()

	fighter1 := NewAircraftManned("F14", "Default", countrycodes.UK, nato.OF2)
	fighter2 := NewAircraftManned("MiG-29", "Default", countrycodes.UK, nato.OF1)

	assert.NotNil(t, fighter1)
	assert.NotNil(t, fighter2)

	ldp1 := DogfightPositionTossup
	ldp2 := DogfightPositionTossup
	dr1, dr2 := Dogfight(&fighter1, ldp1, &fighter2, ldp2)
	assert.NotNil(t, dr1)
	assert.NotNil(t, dr2)
	fmt.Println(dr1)
	fmt.Println(dr2)
}

func TestMoreRounds(t *testing.T) {
	InitGame()

	fighter1 := NewAircraftManned("F14", "Default", countrycodes.UK, nato.OF2)
	fighter2 := NewAircraftManned("MiG-29", "Default", countrycodes.Russia, nato.OF1)

	assert.NotNil(t, fighter1)
	assert.NotNil(t, fighter2)

	drl1, drl2 := Sim10Rounds(&fighter1, &fighter2)
	assert.NotNil(t, drl1)
	assert.NotNil(t, drl2)

	fmt.Println(drl1)
	fmt.Println(drl2)
}
