package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	assert.Nil(t, InitGame(0))
	dfs := CreateSetup()
	fmt.Println(dfs)
}

func TestDeriveSituations(t *testing.T) {
	assert.Nil(t, InitGame(0))
	dfs := CreateSetup()
	fmt.Println(dfs)
}

func TestDogfightSetup_CreateDogfight(t *testing.T) {
	assert.Nil(t, InitGame(0))

	b1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	r1 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)

	ds := NewDogfightSetup()
	ds.AddBlue(b1.AircraftId)
	ds.AddRed(r1.AircraftId)

	assert.Equal(t, b1.AircraftId, ds.TeamBlue[0])
	assert.Equal(t, r1.AircraftId, ds.TeamRed[0])

	d := NewDogfight(ds)

	assert.Equal(t, 0, len(d.Groups))
	assert.Equal(t, b1.AircraftId, d.TeamBlueWaiting[0])
	assert.Equal(t, r1.AircraftId, d.TeamRedWaiting[0])
}

func TestRemoveElementTest(t *testing.T) {
	il := []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, 7, len(il))
	il = append(il[:len(il)-1])
	assert.Equal(t, 6, len(il))
}

func TestAircraftIdListRemoval(t *testing.T) {
	aid1 := AircraftId(uuid.New())
	aid2 := AircraftId(uuid.New())
	aid3 := AircraftId(uuid.New())
	aidl := AircraftIdList([]AircraftId{aid1, aid2, aid3})
	id := aidl.PullFirst()
	assert.Equal(t, aid1, id)
	assert.Equal(t, 2, len(aidl))
	assert.Equal(t, aid2, aidl[0])
	assert.Equal(t, aid3, aidl[1])
}

func TestDogfight_DistributeAircraftsToGroups(t *testing.T) {
	assert.Nil(t, InitGame(0))

	b1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	r1 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
	ds := NewDogfightSetup()
	ds.AddBlue(b1.AircraftId)
	ds.AddRed(r1.AircraftId)
	d := NewDogfight(ds)
	d.DistributeAircraftsToGroups()

	assert.Equal(t, 0, len(d.TeamBlueWaiting))
	assert.Equal(t, 0, len(d.TeamRedWaiting))
	assert.Equal(t, 1, len(d.Groups))
	assert.Equal(t, b1.AircraftId, d.Groups[0].BlueFighterId)
	assert.Equal(t, r1.AircraftId, d.Groups[0].RedFighterId)
}

func TestDogfight_DistributeAircraftsToGroupsMore(t *testing.T) {
	assert.Nil(t, InitGame(0))

	b1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	b2 := NewAircraft("F14", "Default", WarPartyIdUSA)
	b3 := NewAircraft("F14", "Default", WarPartyIdUSA)
	b4 := NewAircraft("F14", "Default", WarPartyIdUSA)
	b5 := NewAircraft("F14", "Default", WarPartyIdUSA)
	r1 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
	r2 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
	ds := NewDogfightSetup()
	ds.AddBlue(b1.AircraftId)
	ds.AddBlue(b2.AircraftId)
	ds.AddBlue(b3.AircraftId)
	ds.AddBlue(b4.AircraftId)
	ds.AddBlue(b5.AircraftId)
	ds.AddRed(r1.AircraftId)
	ds.AddRed(r2.AircraftId)
	d := NewDogfight(ds)
	assert.True(t, d.DistributeAircraftsToGroups())
	assert.False(t, d.DistributeAircraftsToGroups())
	assert.Equal(t, 1, len(d.TeamBlueWaiting))
	assert.Equal(t, b5.AircraftId, d.TeamBlueWaiting[0])
	assert.Equal(t, 0, len(d.TeamRedWaiting))
	assert.Equal(t, 2, len(d.Groups))
	assert.Equal(t, b1.AircraftId, d.Groups[0].BlueFighterId)
	assert.Equal(t, b3.AircraftId, *d.Groups[0].BlueSupportId)
	assert.Equal(t, r1.AircraftId, d.Groups[0].RedFighterId)
	assert.Nil(t, d.Groups[0].RedSupportId)
	assert.Equal(t, b2.AircraftId, d.Groups[1].BlueFighterId)
	assert.Equal(t, b4.AircraftId, *d.Groups[1].BlueSupportId)
	assert.Equal(t, r2.AircraftId, d.Groups[1].RedFighterId)
	assert.Nil(t, d.Groups[1].RedSupportId)
}

func TestDogfight_Simulate(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ds := NewDogfightSetup()
	for i := 0; i < 50; i++ {
		b := NewAircraft("F14", "Default", WarPartyIdUSA)
		bpl := NewPilots(2, WarPartyIdUSA, nato.OF1)
		b.FillSeatsWith(bpl)
		ds.AddBlue(b.AircraftId)
		r := NewAircraft("F14", "Default", WarPartyIdRussia)
		rpl := NewPilots(2, WarPartyIdRussia, nato.OF1)
		r.FillSeatsWith(rpl)
		ds.AddRed(r.AircraftId)
	}
	d := NewDogfight(ds)
	assert.True(t, d.DistributeAircraftsToGroups())
	for round := 0; round < 10; round++ {
		d.Simulate()
	}
}

func TestDistributeReshuffle(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.ErrorLevel))
	Log.SetLevel(logrus.ErrorLevel)
	ds := NewDogfightSetup()

	b1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	bpl1 := NewPilots(2, WarPartyIdUSA, nato.OF1)
	b1.FillSeatsWith(bpl1)
	ds.AddBlue(b1.AircraftId)
	b2 := NewAircraft("F14", "Default", WarPartyIdUSA)
	bpl2 := NewPilots(2, WarPartyIdUSA, nato.OF1)
	b2.FillSeatsWith(bpl2)
	ds.AddBlue(b2.AircraftId)
	b3 := NewAircraft("F14", "Default", WarPartyIdUSA)
	bpl3 := NewPilots(2, WarPartyIdUSA, nato.OF1)
	b3.FillSeatsWith(bpl3)
	ds.AddBlue(b3.AircraftId)

	r1 := NewAircraft("F14", "Default", WarPartyIdRussia)
	rpl1 := NewPilots(2, WarPartyIdRussia, nato.OF1)
	r1.FillSeatsWith(rpl1)
	ds.AddRed(r1.AircraftId)
	r2 := NewAircraft("F14", "Default", WarPartyIdRussia)
	rpl2 := NewPilots(2, WarPartyIdRussia, nato.OF1)
	r2.FillSeatsWith(rpl2)
	ds.AddRed(r2.AircraftId)
	r3 := NewAircraft("F14", "Default", WarPartyIdRussia)
	rpl3 := NewPilots(2, WarPartyIdRussia, nato.OF1)
	r3.FillSeatsWith(rpl3)
	ds.AddRed(r3.AircraftId)

	d := NewDogfight(ds)
	d.DistributeAircraftsToGroups()
	fmt.Println(d.String())
	Globals.AllAircrafts[r1.AircraftId].Destroyed = true
	d.DistributeAircraftsToGroups()
	fmt.Println(d.String())
	assert.Equal(t, 2, len(d.Groups))
	
}
