package bgw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AircraftType uint

const (
	F14   AircraftType = 0
	Mig23              = 1
	Mig27              = 2
	Su17               = 3
)

type Aircraft struct {
	Type               AircraftType
	Altitude           AltitudeBand // Aktuelle Höhe.
	CurrentPosition    Position
	NextTargetLocation Position // Das ist die Position, die das Flugzeug jetzt ansteuert.
	Hitpoints
}

type AircraftParameters struct {
	Type                  AircraftType
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
	Configurations        []WeaponSystemConfiguration
	MaintenanceTime       Rating
	StructuralDefense     Rating
	MaxHitpoints          Hitpoints
}

type AircraftLibrary []AircraftParameters

var AirLib AircraftLibrary

type AircraftParametersNotFound struct {
	Type AircraftType
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

func (a Aircraft) GetParameters() (*AircraftParameters, error) {
	if AirLib == nil {
		return nil, nil
	}
	for _, parameters := range AirLib {
		if parameters.Type == a.Type {
			return &parameters, nil
		}
	}
	return nil, &AircraftParametersNotFound{Type: a.Type}
}

