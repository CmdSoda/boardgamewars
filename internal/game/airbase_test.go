package game

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAirbase(t *testing.T) {
	assert.Nil(t, InitGame())
	ab := NewAirbase("Fallujah AB", &WarPartyIdUSA, Position{3, 5})
	ab.CreateAircrafts("F14", "Default", &WarPartyIdUSA, 6)
	fmt.Println(ab.String())
	fmt.Println(Globals.AirbaseList)
}

func TestHangar(t *testing.T) {
	assert.Nil(t, InitGame())
	ab := NewAirbase("Fallujah AB", &WarPartyIdUSA, Position{3, 5})
	ac := NewAircraft("F14", "Default", &WarPartyIdUSA)
	ab.AddToHangar(ac)
	fmt.Println(ab)
}
