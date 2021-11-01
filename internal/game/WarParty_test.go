package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Nil(t, InitGame())
	wp := NewWarParty("USA", countrycodes.USA, Blue)
	ac := NewAircraft("F14", "Default", wp.WarPartyId)
	pl := NewPilots(ac.GetParameters().Seats, wp.WarPartyId, nato.OF1)
	ac.FillSeatsWith(pl)
	wp.Aircrafts = append(wp.Aircrafts, ac.AircraftId)
	fmt.Println(wp)
}

func TestAddresses(t *testing.T) {
	m := map[int]string{}
	m[0] = "Patrick"
	m[1] = "Uschi"
	k := m
	k[2] = "Peter"
	fmt.Printf("%p\n", &m) // Adresse von m: "0xc000010060"
	fmt.Printf("%p\n", &k) // Adresse von k: "0xc000010068"
	assert.Equal(t, "Peter", m[2])
	assert.Equal(t, "Peter", k[2])
	delete(k, 1)
	_, mExist := m[1]
	_, kExist := m[1]
	assert.Equal(t, false, mExist)
	assert.Equal(t, false, kExist)
}
