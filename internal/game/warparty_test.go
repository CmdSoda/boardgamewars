package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Nil(t, InitGame(0))
	wp := NewWarParty("usa", Blue)
	ac := NewAircraft("F14", "Default", wp.WarPartyId)
	pl := NewPilots(ac.GetParameters().Seats, wp.WarPartyId, OF1)
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
	assert.Nil(t, InitGame(0))
	wp := NewWarParty("usa", Blue)
	fmt.Println(wp)
	ab := NewAirbase("Parkhaus", wp.WarPartyId, hexagon.HexPosition{Column: 1, Row: 1})
	assert.Equal(t, CountryName("usa"), Globals.AllWarParties[wp.WarPartyId].CountryName)
	ab.AssignToWarParty(wp.WarPartyId)
	assert.Equal(t, wp.WarPartyId, ab.BelongsTo)
	//assert.Equal(t, wp.Airbases)
}

func TestBlueRed(t *testing.T) {
	assert.Nil(t, InitGame(0))
	count := len(Globals.AllWarParties)
	wp1 := NewWarParty("usa2", Blue)
	wp2 := NewWarParty("russia2", Red)
	wp3 := NewWarParty("russia3", 99)
	fmt.Println(wp1)
	fmt.Println(wp2)
	fmt.Println(wp3)
	assert.Equal(t, count+3, len(Globals.AllWarParties))
}

func subfunction(il []int) {
	il[0] = 99
}

func TestSlice2Evil(t *testing.T) {
	l := make([]int, 1) // "{ 0 }"
	subfunction(l)      // "{ 99 }"
	assert.Equal(t, 99, l[0])
}

func TestLoadWarParties(t *testing.T) {
	assert.Nil(t, InitGame(0))
	assert.Equal(t, CountryName("uk"), Globals.AllWarParties[WarPartyIdUK].CountryName)
	assert.Equal(t, CountryName("germany"), Globals.AllWarParties[WarPartyIdGermany].CountryName)
	assert.Equal(t, CountryName("russia"), Globals.AllWarParties[WarPartyIdRussia].CountryName)
	assert.Equal(t, CountryName("usa"), Globals.AllWarParties[WarPartyIdUSA].CountryName)
}
