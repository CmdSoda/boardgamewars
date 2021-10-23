package bgw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AircraftId int

type Aircraft struct {
	AircraftId
	Altitude           AltitudeBand // Aktuelle Höhe.
	CurrentPosition    Position
	NextTargetLocation Position // Das ist die Position, die das Flugzeug jetzt ansteuert.
	WeaponSystems      []WeaponSystem
	Hitpoints
}

func NewAircraftById(id AircraftId, configurationName string) *Aircraft {
	ac := Aircraft{AircraftId: id}
	ac.WeaponSystems = NewWeaponSystems(id, configurationName)
	for i := 0; i < len(ac.WeaponSystems); i++ {
		ac.WeaponSystems[i].InitWeaponSystem()
	}
	return &ac
}

func NewAircraftByName(name string, configurationName string) *Aircraft {
	id := GetAircraftIdByName(name)
	if id >= 0 {
		return NewAircraftById(id, configurationName)
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
	var max int = 0
	for _, system := range a.WeaponSystems {
		if system.Air2AirWeaponParameters != nil {
			if int(system.Air2AirWeaponParameters.Dogfighting) > max {
				bestws = &system
				max = int(system.Dogfighting)
			}
		}
	}
	return bestws
}