package game

import (
	"encoding/json"
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"strings"
)

type AircraftId int

type Aircraft struct {
	AircraftId
	countrycodes.Code
	Id                 uuid.UUID
	Altitude           AltitudeBand // Aktuelle Höhe.
	CurrentPosition    Position
	NextTargetLocation Position // Das ist die Position, die das Flugzeug jetzt ansteuert.
	WeaponSystems      []WeaponSystem
	Damage             []DamageType // Eine Liste von Schäden
	Destroyed          bool
	Pilots             []uuid.UUID
	StationedAt        uuid.UUID
}

func (a Aircraft) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s\nPilots: ", a.GetParameters().Name)
	for _, pilot := range a.Pilots {
		fmt.Fprintf(&b, TheRoster.GetPilot(pilot).String()+" ")
	}
	fmt.Fprint(&b, "\nDamage: ")
	for _, d := range a.Damage {
		fmt.Fprintf(&b, d.String()+" ")
	}
	return b.String()
}

func (a *Aircraft) AddPilot(p Pilot) {
	a.Pilots = append(a.Pilots, p.UUID)
	TheRoster.Add(p)
}

func (a *Aircraft) FillUpSeats(oc nato.Code) {
	a.Pilots = make([]uuid.UUID, 0)
	currentoc := oc
	for i := 0; i < a.GetParameters().Seats; i++ {
		a.AddPilot(NewPilot(a.Code, currentoc))
		if currentoc > 1 {
			currentoc = currentoc - 1
		}
	}
}

func (a *Aircraft) AssignToAB(id uuid.UUID) bool {
	_, exist := AllAirbases[id]
	if exist {
		a.StationedAt = id
		return true
	}
	return false
}

func NewAircraftManned(name string, configurationName string, cc countrycodes.Code, oc nato.Code) *Aircraft {
	ac := NewAircraft(name, configurationName, cc)
	ac.FillUpSeats(oc)
	return ac
}

func NewAircraft(name string, configurationName string, cc countrycodes.Code) *Aircraft {
	id := GetAircraftIdByName(name)
	if id >= 0 {
		ac := Aircraft{AircraftId: id}
		ac.Id = uuid.New()
		ac.Code = cc
		ac.WeaponSystems = NewWeaponSystems(id, configurationName)
		for i := 0; i < len(ac.WeaponSystems); i++ {
			ac.WeaponSystems[i].InitWeaponSystem()
		}
		ac.Damage = make([]DamageType, 0)
		return &ac
	}
	return nil
}

func GetAircraftIdByName(name string) AircraftId {
	for _, parameters := range AirLib {
		if parameters.Name == name {
			return parameters.AircraftId
		}
	}
	return -1
}

type AircraftParameters struct {
	AircraftId
	Name                  string
	Nickname              string
	FirstFlight           Year
	Introduction          Year
	CombatSpeed           Rating
	CruiseSpeed           Rating
	CombatFuelConsumption Rating // Treibstoffverbrauch im Kampf pro Runde.
	CruiseFuelConsumption Rating // Treibstoffverbrauch beim Cruisen pro Runde.
	Fuel                  Rating
	MaxAltitude           AltitudeBand
	Dogfighting           Rating
	Configurations        WeaponSystemConfigurationList
	MaintenanceTime       Rating
	StructuralDefense     Rating
	MaxHitpoints          Hitpoints
	MaxDamagePoints       int
	Seats                 int
}

type AircraftLibrary []AircraftParameters

var AirLib AircraftLibrary

type AircraftParametersNotFound struct {
	Type AircraftId
}

func (p *AircraftParametersNotFound) Error() string {
	return fmt.Sprintf("Could not find parameters for aircraft %d", p.Type)
}

func LoadAircrafts() (*AircraftLibrary, error) {
	var err error
	file, err := os.Open("data/aircrafts.json")
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	al := AircraftLibrary{}
	err = json.Unmarshal(bytes, &al)
	if err != nil {
		return nil, err
	}
	AirLib = al
	return &al, nil
}

func (a Aircraft) GetParameters() *AircraftParameters {
	for _, parameters := range AirLib {
		if parameters.AircraftId == a.AircraftId {
			return &parameters
		}
	}
	return nil
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
				ws.Air2AirWeaponParameters.EquipmentId == system.Air2AirWeaponParameters.EquipmentId {
				a.WeaponSystems[i].Depleted = true
				return
			}
		}
	}
}
