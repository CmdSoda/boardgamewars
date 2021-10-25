package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/google/uuid"
	"strings"
)

// An airbase can be at land and on the carrier

type Airbase struct {
	Id           uuid.UUID
	Name         string
	BelongsTo    countrycodes.Code
	AcceptAllies bool
	Aircrafts    []Aircraft
	Position
}

func (a Airbase) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Airbase: %s [%s]\n", a.Name, a.BelongsTo.String())
	fmt.Fprint(&sb, "Aircrafts: ")
	for _, aircraft := range a.Aircrafts {
		fmt.Fprintf(&sb, "%s, ", aircraft.GetParameters().Name)
	}
	fmt.Fprint(&sb, "\n")
	fmt.Fprintf(&sb, "Location: %s\n", a.Position)

	return sb.String()
}

func NewAirbase(name string, cc countrycodes.Code, pos Position) Airbase {
	ab := Airbase{}
	ab.Id = uuid.New()
	ab.Name = name
	ab.BelongsTo = cc
	ab.Position = pos
	return ab
}

func (a *Airbase) CreateAircrafts(aircraftName string, configurationName string, cc countrycodes.Code, count int) {
	for i := 0; i < count; i++ {
		ac := NewAircraft(aircraftName, configurationName, cc)
		if ac != nil {
			a.Aircrafts = append(a.Aircrafts, *ac)
		}
	}
}
