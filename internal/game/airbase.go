package game

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// An airbase can be at land and on the carrier

type AirbaseId uuid.UUID

type Airbase struct {
	AirbaseId
	Name                string
	BelongsTo           WarPartyId
	AcceptAllies        bool
	AllAircrafts        []AircraftId // Alle Aircrafts, die zu dieser Basis gehören
	AllPilots           []PilotId    // Alle Piloten, die zu dieser Basis gehören
	AircraftsHangar     []AircraftId
	AircraftsMaintained []AircraftId
	FloatPosition
}

type AirbasesMap map[AirbaseId]Airbase

func (al AirbasesMap) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "Airbases:\n")
	for _, airbase := range al {
		fmt.Fprintf(&sb, "%s\n", airbase.Name)
	}
	return sb.String()
}

func (ab Airbase) String() string {
	var sb strings.Builder
	wp := Globals.AllWarParties[ab.BelongsTo]
	fmt.Fprintf(&sb, "Airbase: %s [%s]\n", ab.Name, wp.String())
	fmt.Fprint(&sb, "AircraftsHangar: ")
	for _, aircraftid := range ab.AircraftsHangar {
		fmt.Fprintf(&sb, "%s, ", Globals.AllAircrafts[aircraftid].GetParameters().Name)
	}
	fmt.Fprint(&sb, "\n")
	fmt.Fprintf(&sb, "Location: %s\n", ab.FloatPosition)

	return sb.String()
}

func NewAirbase(name string, warpartyid WarPartyId, pos FloatPosition) Airbase {
	ab := Airbase{}
	ab.AirbaseId = AirbaseId(uuid.New())
	ab.Name = name
	ab.BelongsTo = warpartyid
	ab.FloatPosition = pos
	Globals.AllAirbases[ab.AirbaseId] = ab
	ab.AircraftsHangar = []AircraftId{}
	ab.AircraftsMaintained = []AircraftId{}
	ab.AllPilots = []PilotId{}
	ab.AllAircrafts = []AircraftId{}
	return ab
}

func (ab *Airbase) AddToHangar(acid AircraftId) {
	ab.AircraftsHangar = append(ab.AircraftsHangar, acid)
}

func (ab *Airbase) CreateAircrafts(aircraftName string, configurationName string, warpartyid WarPartyId, count int) {
	for i := 0; i < count; i++ {
		ac := NewAircraft(aircraftName, configurationName, warpartyid)
		ac.StationedAt = ab.AirbaseId
		ab.AircraftsHangar = append(ab.AircraftsHangar, ac.AircraftId)
	}
}

func (ab *Airbase) AssignToWarParty(wpid WarPartyId) {
	wp := Globals.AllWarParties[wpid]
	ab.BelongsTo = wpid
	wp.Airbases[ab.AirbaseId] = struct{}{}
}
