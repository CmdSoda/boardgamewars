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

func TestMapReferenceType(t *testing.T) {
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

func TestSlice(t *testing.T) {
	var l []int
	l = append(l, 5)
	k := l
	k = append(k, 7)
	assert.Equal(t, 1, len(l))
	assert.Equal(t, 2, len(k))
}

func TestList(t *testing.T) {
	var x = make([]int, 0)
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	var y = make([]int, 3)
	copy(y, x)
	assert.Equal(t, 3, len(x))
	assert.Equal(t, 3, len(y))
	x = append(x, 4)
	assert.Equal(t, 4, len(x))
	assert.Equal(t, 4, x[3])
	assert.Equal(t, 3, len(y))
	y = append(y, 5)
	assert.Equal(t, 4, len(x))
	assert.Equal(t, 4, x[3])
	assert.Equal(t, 4, len(y))
	assert.Equal(t, 5, y[3])
}

func TestNewWarParty(t *testing.T) {
	assert.Nil(t, InitGame())
	wp := NewWarParty("NeueParty", countrycodes.USA, Blue)
	ab := NewAirbase("Parkhaus", wp.WarPartyId, Position{1, 1})
	assert.Equal(t, wp.Name, Globals.AllWarParties[wp.WarPartyId].Name)
	ab.AssignToWarParty(wp.WarPartyId)
	assert.Equal(t, wp.WarPartyId, ab.BelongsTo)
	//assert.Equal(t, wp.Airbases)
}

