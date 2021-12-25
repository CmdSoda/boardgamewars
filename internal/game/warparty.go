package game

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)

type WarPartyId uuid.UUID

type WarPartyColor uint

const (
	Blue WarPartyColor = 0
	Red  WarPartyColor = 1
)

func (wps WarPartyColor) String() string {
	switch wps {
	case Blue:
		return "Blue"
	case Red:
		return "Red"
	}
	return "Unknown"
}

func WarPartyColorFromString(wpc string) WarPartyColor {
	switch wpc {
	case "blue":
		return Blue
	case "red":
		return Red
	}
	return Blue
}

type WarParty struct {
	WarPartyId
	WarPartyIdString    string
	WarPartyColorString string
	WarPartyColor
	CountryName
	Pilots    []PilotId
	Aircrafts []AircraftId
	Airbases  map[AirbaseId]struct{}
}

func (w WarParty) String() string {
	return fmt.Sprintf("%s [%s]\nAircrafts: %d", w.CountryName, w.WarPartyColor.String(), len(w.Aircrafts))
}

type WarPartyMap map[WarPartyId]*WarParty

func NewWarParty(cn CountryName, side WarPartyColor) WarParty {
	wp := WarParty{}
	wp.WarPartyColor = side
	wp.WarPartyId = WarPartyId(uuid.New())
	wp.Pilots = []PilotId{}
	wp.Aircrafts = []AircraftId{}
	wp.Airbases = map[AirbaseId]struct{}{}
	wp.CountryName = cn
	Globals.AllWarParties[wp.WarPartyId] = &wp
	return wp
}

func LoadWarParties() (map[WarPartyId]*WarParty, error) {
	wm := map[WarPartyId]*WarParty{}
	dataPathFilename := Globals.Startup.DataPath + "warparties.json"
	file, err := os.Open(dataPathFilename)
	if err != nil {
		Log.Errorf("%s not found\n", dataPathFilename)
		return wm, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		Log.Errorf("%s error while reading\n", dataPathFilename)
		return wm, err
	}
	var wl []WarParty
	err = json.Unmarshal(bytes, &wl)
	if err != nil {
		Log.Errorf("%s error while unmarshaling\n", dataPathFilename)
		return wm, err
	}

	for idx, _ := range wl {
		party := &wl[idx]
		party.WarPartyId = WarPartyId(uuid.MustParse(party.WarPartyIdString))
		party.WarPartyColor = WarPartyColorFromString(party.WarPartyColorString)
		wm[party.WarPartyId] = party
	}

	return wm, nil
}
