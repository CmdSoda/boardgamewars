package game

import (
	"encoding/json"
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
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
	Country       countrycodes.Code
	CountryString string
	Pilots        []PilotId
	Aircrafts     []AircraftId
	Airbases      map[AirbaseId]struct{}
}

func (w WarParty) String() string {
	return fmt.Sprintf("%s [%s]\nAircrafts: %d", w.Country.String(), w.WarPartyColor.String(), len(w.Aircrafts))
}

type WarPartyMap map[WarPartyId]*WarParty

func NewWarParty(code countrycodes.Code, side WarPartyColor) WarParty {
	wp := WarParty{}
	wp.Country = code
	wp.WarPartyColor = side
	wp.WarPartyId = WarPartyId(uuid.New())
	wp.Pilots = []PilotId{}
	wp.Aircrafts = []AircraftId{}
	wp.Airbases = map[AirbaseId]struct{}{}
	Globals.AllWarParties[wp.WarPartyId] = &wp
	return wp
}

func LoadWarParties(filename string) (map[WarPartyId]*WarParty, error) {
	wm := map[WarPartyId]*WarParty{}
	file, err := os.Open(filename)
	if err != nil {
		Log.Errorf("%s not found\n", filename)
		return wm, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		Log.Errorf("%s error while reading\n", filename)
		return wm, err
	}
	var wl []WarParty
	err = json.Unmarshal(bytes, &wl)
	if err != nil {
		Log.Errorf("%s error while unmarshaling\n", filename)
		return wm, err
	}

	for _, party := range wl {
		party.WarPartyId = WarPartyId(uuid.MustParse(party.WarPartyIdString))
		party.WarPartyColor = WarPartyColorFromString(party.WarPartyColorString)
		wm[party.WarPartyId] = &party
	}

	return wm, nil
}
