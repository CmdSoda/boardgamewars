package bgw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func CreateAircraftLibrary() {
	ap := AircraftParameters{}
	ap.Type = F14
	ap.Name = "F14"
	ap.FirstFlight = 1970
	ap.Introduction = 1974
	ap.CombatSpeed = 10
	ap.CruiseSpeed = 4
	ap.CombatFuelConsumption = 10
	ap.CombatFuelConsumption = 4
	ap.Fuel = 20
	ap.MaxAltitude = 5
	ap.Dogfighting = 7
	wsc := WeaponSystemConfiguration{FuelTank, FuelTank}
	ap.Configurations = []WeaponSystemConfiguration{wsc}
	ap.MaintenanceTime = 5
	al := AircraftLibrary{ap}
	jb, _ := json.Marshal(al)
	fmt.Println(string(jb))
}

func TestLoadAircraftLibrary(t *testing.T) {
	var err error
	file, err := os.Open("data/aircrafts.json")
	if err != nil {
		t.FailNow()
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.FailNow()
	}
	al := AircraftLibrary{}
	err = json.Unmarshal(bytes, &al)
	if err != nil {
		t.FailNow()
	}

	if len(al) == 0 {
		t.FailNow()
	}
}