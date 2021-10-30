package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/google/uuid"
)

type WarPartyId uuid.UUID

type WarPartyColor uint

const (
	Blue WarPartyColor = 0
	Red  WarPartyColor = 1
)

func (wps *WarPartyColor) String() string {
	switch *wps {
	case Blue:
		return "Blue"
	case Red:
		return "Red"
	}
	return "Unknown"
}

type WarParty struct {
	WarPartyId
	WarPartyColor
	Name      string
	Country   countrycodes.Code
	Pilots    map[PilotId]*Pilot
	Aircrafts map[AircraftId]*Aircraft
}

func (w *WarParty) String() string {
	return fmt.Sprintf("%s [%s]\nAircrafts: %d", w.Name, w.WarPartyColor.String(), len(w.Aircrafts))
}

type WarPartyList map[WarPartyId]WarParty

func NewWarParty(name string, code countrycodes.Code, side WarPartyColor) *WarParty {
	wp := WarParty{}
	wp.Name = name
	wp.Country = code
	wp.WarPartyColor = side
	wp.WarPartyId = WarPartyId(uuid.New())
	wp.Pilots = map[PilotId]*Pilot{}
	wp.Aircrafts = map[AircraftId]*Aircraft{}
	Globals.WarPartyList[wp.WarPartyId] = wp
	return &wp
}
