package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/google/uuid"
	"strings"
)

// An airbase can be at land and on the carrier

type Airbase struct {
	Id                  uuid.UUID
	Name                string
	BelongsTo           countrycodes.Code
	AcceptAllies        bool
	AircraftsHangar     []Aircraft
	AircraftsMaintained []Aircraft
	StationedPilots     []Pilot
	Position
}

type AirbaseList map[uuid.UUID]Airbase

var AllAirbases AirbaseList

func NewAirbaseList() {
	AllAirbases = map[uuid.UUID]Airbase{}
}

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
	fmt.Fprintf(&sb, "Airbase: %s [%s]\n", ab.Name, ab.BelongsTo.String())
	fmt.Fprint(&sb, "AircraftsHangar: ")
	for _, aircraft := range ab.AircraftsHangar {
		fmt.Fprintf(&sb, "%s, ", aircraft.GetParameters().Name)
	}
	fmt.Fprint(&sb, "\n")
	fmt.Fprintf(&sb, "Location: %s\n", ab.Position)

	return sb.String()
}

func NewAirbase(name string, cc countrycodes.Code, pos Position) Airbase {
	ab := Airbase{}
	ab.Id = uuid.New()
	ab.Name = name
	ab.BelongsTo = cc
	ab.Position = pos
	AllAirbases[ab.Id] = ab
	ab.AircraftsHangar = []Aircraft{}
	ab.StationedPilots = []Pilot{}
	return ab
}

func (ab *Airbase) AddToHangar(ac Aircraft) {
	ab.AircraftsHangar = append(ab.AircraftsHangar, ac)
}

func (ab *Airbase) CreateAircrafts(aircraftName string, configurationName string, cc countrycodes.Code, count int) {
	for i := 0; i < count; i++ {
		ac := NewAircraft(aircraftName, configurationName, cc)
		ac.StationedAt = ab.Id
		if ac != nil {
			ab.AircraftsHangar = append(ab.AircraftsHangar, *ac)
		}
	}
}
