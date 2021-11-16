package game

import (
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)

type AircraftParametersId uuid.UUID

var InvalidAircraftParametersId = AircraftParametersId(uuid.MustParse("24b30695-6da6-4037-932c-f413e06c5ade"))

type AircraftParameters struct {
	Id                    AircraftParametersId
	IdAsString            string
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

type AircraftParametersMap map[AircraftParametersId]AircraftParameters

type AircraftLibraryFile []AircraftParameters

func LoadAircraftParameters() error {
	var err error
	filename := Globals.Configuration.DataPath + "/aircraftparameters.json"
	file, err := os.Open(filename)
	if err != nil {
		Log.Errorf("%s not found\n", filename)
		return err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		Log.Errorf("%s while reading\n", filename)
		return err
	}
	al := AircraftLibraryFile{}
	err = json.Unmarshal(bytes, &al)
	if err != nil {
		Log.Errorf("%s while unmarshaling\n", filename)
		return err
	}
	Globals.AllAircraftParameters = AircraftParametersMap{}
	for _, parameters := range al {
		parameters.Id = AircraftParametersId(uuid.MustParse(parameters.IdAsString))
		Globals.AllAircraftParameters[parameters.Id] = parameters
	}
	return nil
}
