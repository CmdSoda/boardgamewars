package game

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// An airbase can be at land and on the carrier

type AirbaseId uuid.UUID

type Airbase struct {
	Id                  AirbaseId
	Name                string
	BelongsTo           *WarPartyId
	AcceptAllies        bool
	AircraftsHangar     map[AircraftId]*Aircraft
	AircraftsMaintained map[AircraftId]*Aircraft
	StationedPilots     map[PilotId]*Pilot
	Position
}

type AirbaseList map[AirbaseId]Airbase

func (al AirbaseList) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "Airbases:\n")
	for _, airbase := range al {
		fmt.Fprintf(&sb, "%s\n", airbase.Name)
	}
	return sb.String()
}

func (ab Airbase) String() string {
	var sb strings.Builder
	wp := Globals.WarPartyList[*ab.BelongsTo]
	fmt.Fprintf(&sb, "Airbase: %s [%s]\n", ab.Name, wp.String())
	fmt.Fprint(&sb, "AircraftsHangar: ")
	for _, aircraft := range ab.AircraftsHangar {
		fmt.Fprintf(&sb, "%s, ", aircraft.GetParameters().Name)
	}
	fmt.Fprint(&sb, "\n")
	fmt.Fprintf(&sb, "Location: %s\n", ab.Position)

	return sb.String()
}

func NewAirbase(name string, warpartyid *WarPartyId, pos Position) Airbase {
	ab := Airbase{}
	ab.Id = AirbaseId(uuid.New())
	ab.Name = name
	ab.BelongsTo = warpartyid
	ab.Position = pos
	Globals.AirbaseList[ab.Id] = ab
	ab.AircraftsHangar = map[AircraftId]*Aircraft{}
	ab.AircraftsMaintained = map[AircraftId]*Aircraft{}
	ab.StationedPilots = map[PilotId]*Pilot{}
	return ab
}

func (ab *Airbase) AddToHangar(ac *Aircraft) {
	ab.AircraftsHangar[ac.AircraftId] = ac
}

func (ab *Airbase) CreateAircrafts(aircraftName string, configurationName string, warpartyid *WarPartyId, count int) {
	for i := 0; i < count; i++ {
		ac := NewAircraft(aircraftName, configurationName, warpartyid)
		ac.StationedAt = &ab.Id
		ab.AircraftsHangar[ac.AircraftId] = ac
	}
}
