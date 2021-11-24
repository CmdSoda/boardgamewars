package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/google/uuid"
	"strings"
)

// An airbase can be at land and on the carrier

type AirbaseId uuid.UUID

type Airbase struct {
	AirbaseId
	Name                   string
	BelongsTo              WarPartyId
	AcceptAllies           bool
	AllAircrafts           []AircraftId // Alle Aircrafts, die zu dieser Basis gehören
	AllPilots              []PilotId    // Alle Piloten, die zu dieser Basis gehören
	ParkingArea            []AircraftId
	MaxParkingSlots        int
	MaintenanceArea        []AircraftId
	MaxMaintainenanceSlots int
	hexagon.HexPosition
}

type AirbasesMap map[AirbaseId]Airbase

//goland:noinspection GoUnhandledErrorResult
func (al AirbasesMap) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "Airbases:\n")
	for _, airbase := range al {
		fmt.Fprintf(&sb, "%s\n", airbase.Name)
	}
	return sb.String()
}

func NewAirbase(name string, warpartyid WarPartyId, pos hexagon.HexPosition) *Airbase {
	ab := Airbase{}
	ab.AirbaseId = AirbaseId(uuid.New())
	ab.Name = name
	ab.BelongsTo = warpartyid
	ab.HexPosition = pos
	ab.ParkingArea = []AircraftId{}
	ab.MaintenanceArea = []AircraftId{}
	ab.AllPilots = []PilotId{}
	ab.AllAircrafts = []AircraftId{}
	Globals.AllAirbases[ab.AirbaseId] = ab
	return &ab
}

func (ab *Airbase) AddToParkingArea(acid AircraftId) {
	ab.ParkingArea = append(ab.ParkingArea, acid)
}

func (ab *Airbase) CreateAircrafts(aircraftName string, configurationName string, warpartyid WarPartyId, count int) {
	for i := 0; i < count; i++ {
		ac := NewAircraft(aircraftName, configurationName, warpartyid)
		ac.StationedAt = ab.AirbaseId
		ab.ParkingArea = append(ab.ParkingArea, ac.AircraftId)
	}
}

func (ab *Airbase) AssignToWarParty(wpid WarPartyId) {
	wp := Globals.AllWarParties[wpid]
	ab.BelongsTo = wpid
	wp.Airbases[ab.AirbaseId] = struct{}{}
}

// moveAircraftToMaintenance bewegt ein Flugzeug von ParkingArea zu MaintenanceArea.
func (ab *Airbase) moveAircraftToMaintenance(idx int) {
	// Aircraft aus ParkingArea entfernen und in
	acid := ab.ParkingArea[idx]
	ab.ParkingArea = append(ab.ParkingArea[:idx], ab.ParkingArea[idx+1:]...)
	ab.MaintenanceArea = append(ab.MaintenanceArea, acid)
}

// moveAircraftToParkingArea bewegt ein Flugzeug von MaintenanceArea zu ParkingArea.
func (ab *Airbase) moveAircraftToParkingArea(idx int) {
	acid := ab.MaintenanceArea[idx]
	ab.MaintenanceArea = append(ab.MaintenanceArea[:idx], ab.MaintenanceArea[idx+1:]...)
	ab.ParkingArea = append(ab.ParkingArea, acid)
}

func (ab *Airbase) CalculateRepairTime(ac *Aircraft) {
	ac.RepairTime = 0
	for _, damageType := range ac.Damage {
		mp := Globals.Settings.DamageMaintenanceMultiplier[damageType.String()]
		cost := mp * Globals.Settings.RepairTimePerDamageTypeBase
		ac.RepairTime = ac.RepairTime + StepTime(cost)
	}
}

func (ab *Airbase) Step(st StepTime) {
	doAgain := true

	// Flugzeuge im Wartungsbereich entsprechend der StepTime reparieren.
	for doAgain {
		doAgain = false
		for i := range ab.MaintenanceArea {
			// Beschädigtes Flugzeug?
			ac := Globals.AllAircrafts[ab.MaintenanceArea[i]]
			ac.RepairTime = ac.RepairTime - st
			if ac.RepairTime < 0 {
				ac.RepairTime = 0
				ac.Damage = []DamageType{}
				ab.moveAircraftToParkingArea(i)
				doAgain = true
				break
			}
		}
	}

	// Defekte Flugzeuge in den Wartungsbereich schieben
	doAgain = true
	for doAgain {
		doAgain = false
		// Ist noch Platz für mehr Flugzeuge in der Wartung?
		if ab.MaxMaintainenanceSlots > len(ab.MaintenanceArea) {
			for i := range ab.ParkingArea {
				// Beschädigtes Flugzeug?
				ac := Globals.AllAircrafts[ab.ParkingArea[i]]
				if ac.IsDamaged() {
					ab.CalculateRepairTime(ac)
					ab.moveAircraftToMaintenance(i)
					doAgain = true
					break
				}
			}
		}
	}
}

//goland:noinspection GoUnhandledErrorResult
func (ab Airbase) String() string {
	var sb strings.Builder
	wp := Globals.AllWarParties[ab.BelongsTo]
	fmt.Fprintf(&sb, "Airbase: %s [%s]\n", ab.Name, wp.String())
	fmt.Fprint(&sb, "ParkingArea: ")
	for _, aircraftid := range ab.ParkingArea {
		fmt.Fprintf(&sb, "%s, ", Globals.AllAircrafts[aircraftid].GetParameters().Name)
	}
	fmt.Fprint(&sb, "\n")
	fmt.Fprintf(&sb, "Location: %s\n", ab.HexPosition)

	return sb.String()
}
