package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Nil(t, InitGame(0))
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
	assert.Nil(t, InitGame(0))
	wp := NewWarParty("NeueParty", countrycodes.USA, Blue)
	fmt.Println(wp)
	ab := NewAirbase("Parkhaus", wp.WarPartyId, FloatPosition{1, 1})
	assert.Equal(t, wp.Name, Globals.AllWarParties[wp.WarPartyId].Name)
	ab.AssignToWarParty(wp.WarPartyId)
	assert.Equal(t, wp.WarPartyId, ab.BelongsTo)
	//assert.Equal(t, wp.Airbases)
}

func TestBlueRed(t *testing.T) {
	assert.Nil(t, InitGame(0))
	wp1 := NewWarParty("NeueParty", countrycodes.USA, Blue)
	wp2 := NewWarParty("NeueParty2", countrycodes.Russia, Red)
	wp3 := NewWarParty("NeueParty3", countrycodes.Russia, 99)
	fmt.Println(wp1)
	fmt.Println(wp2)
	fmt.Println(wp3)
}

func subfunction(il []int) {
	il[0] = 99
}

func TestSlice2Evil(t *testing.T) {
	l := make([]int, 1) // "{ 0 }"
	subfunction(l)      // "{ 99 }"
	assert.Equal(t, 99, l[0])
}
