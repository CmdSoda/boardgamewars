package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/google/uuid"
	"strings"
)

type AircraftId uuid.UUID

var InvalidAircraftId = AircraftId(uuid.MustParse("0a491791-3cd8-4316-bacf-de84f5e8df27"))

type Aircraft struct {
	AircraftId
	AircraftParametersId
	Country            countrycodes.Code
	Altitude           AltitudeBand // Aktuelle Höhe.
	CurrentPosition    Position
	NextTargetLocation Position // Das ist die Position, die das Flugzeug jetzt ansteuert.
	WeaponSystems      []WeaponSystem
	Damage             []DamageType // Eine Liste von Schäden
	Destroyed          bool
	Pilots             []PilotId
	StationedAt        uuid.UUID
}

func (a Aircraft) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s\nPilots: ", a.GetParameters().Name)
	for _, pilot := range a.Pilots {
		fmt.Fprintf(&b, Globals.PilotRoster.GetPilot(pilot).String()+" ")
	}
	fmt.Fprint(&b, "\nDamage: ")
	for _, d := range a.Damage {
		fmt.Fprintf(&b, d.String()+" ")
	}
	if len(a.Damage) == 0 {
		fmt.Fprint(&b, "\n")
	}
	return b.String()
}

func (a *Aircraft) AddPilot(p Pilot) {
	a.Pilots = append(a.Pilots, p.PilotId)
	Globals.PilotRoster.Add(p)
}

func (a *Aircraft) FillUpSeats(oc nato.Code) {
	a.Pilots = make([]PilotId, 0)
	currentoc := oc

	for i := 0; i < a.GetParameters().Seats; i++ {
		a.AddPilot(NewPilot(a.Country, currentoc))
		if currentoc > 1 {
			currentoc = currentoc - 1
		}
	}
}

func (a *Aircraft) AssignToAB(id uuid.UUID) bool {
	_, exist := Globals.AirbaseList[id]
	if exist {
		a.StationedAt = id
		return true
	}
	return false
}

func NewAircraftManned(name string, configurationName string, cc countrycodes.Code, oc nato.Code) *Aircraft {
	ac := NewAircraft(name, configurationName, cc)
	if ac.AircraftId != InvalidAircraftId {
		ac.FillUpSeats(oc)
	}
	return ac
}

func NewAircraft(name string, configurationName string, cc countrycodes.Code) *Aircraft {
	ac := Aircraft{}
	acpid := GetAircraftParametersIdByName(name)
	if acpid != InvalidAircraftParametersId {
		ac.AircraftId = AircraftId(uuid.New())
		ac.AircraftParametersId = acpid
		ac.Country = cc
		ac.WeaponSystems = NewWeaponSystems(acpid, configurationName)
		for i := 0; i < len(ac.WeaponSystems); i++ {
			ac.WeaponSystems[i].InitWeaponSystem()
		}
		ac.Damage = make([]DamageType, 0)
		return &ac
	}
	return &ac
}

func GetAircraftParametersIdByName(name string) AircraftParametersId {
	for _, parameters := range Globals.AircraftLibrary {
		if parameters.Name == name {
			return parameters.Id
		}
	}
	panic("unkown aircraft name: " + name)
	return InvalidAircraftParametersId
}

func (a Aircraft) GetParameters() AircraftParameters {
	ap, _ := Globals.AircraftLibrary[a.AircraftParametersId]
	return ap
}

func (a Aircraft) GetBestDogfightingWeapon() *WeaponSystem {
	var bestws *WeaponSystem = nil
	var max = 0
	for _, system := range a.WeaponSystems {
		if system.Depleted == false && system.Air2AirWeaponParameters != nil {
			if int(system.Air2AirWeaponParameters.Dogfighting) > max {
				bestws = &system
				max = int(system.Dogfighting)
			}
		}
	}
	return bestws
}

func (a *Aircraft) AddLightDamage(dt DamageType) {
	a.Damage = append(a.Damage, dt)
}

func (a *Aircraft) DoDamageAssessment() {
	if len(a.Damage) > a.GetParameters().MaxDamagePoints {
		a.Destroyed = true
	}

	// Falls die Kanzel getroffen wurde => Pilot tot?
}

func (a *Aircraft) DoDamageWith(ws WeaponSystem) DamageType {
	if ws.Air2AirWeaponParameters != nil {
		dhp := ws.Air2AirWeaponParameters.DoRandomDamage()
		if dhp <= a.GetParameters().MaxHitpoints {
			rd := RollRandomDamage()
			a.AddLightDamage(rd)
			return rd
		}
	}
	return DamageTypeNothing
}

func (a *Aircraft) DepleteWeapon(ws WeaponSystem) {
	if ws.Air2AirWeaponParameters != nil {
		for i, system := range a.WeaponSystems {
			if system.Depleted == false && system.Air2AirWeaponParameters != nil &&
				ws.Air2AirWeaponParameters.Id == system.Air2AirWeaponParameters.Id {
				a.WeaponSystems[i].Depleted = true
				return
			}
		}
	}
}
