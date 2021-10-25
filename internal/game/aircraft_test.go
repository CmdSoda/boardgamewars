package game

import (
	"encoding/json"
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateAircraftLibrary(t *testing.T) {
	ap := AircraftParameters{}
	ap.AircraftId = 0
	ap.Name = "F14"
	ap.Nickname = "Tomcat"
	ap.FirstFlight = 1970
	ap.Introduction = 1974
	ap.CombatSpeed = 10
	ap.CruiseSpeed = 4
	ap.CombatFuelConsumption = 10
	ap.CombatFuelConsumption = 4
	ap.Fuel = 20
	ap.MaxAltitude = 5
	ap.Dogfighting = 7
	wsc := WeaponSystemConfiguration{
		ConfigurationName: "Default",
		WeaponSystems: []WeaponSystem{{
			WeaponSystemName: "Tank500",
			Category:         "DropTank",
		}},
	}
	ap.Configurations = []WeaponSystemConfiguration{wsc}
	ap.MaintenanceTime = 5
	al := AircraftLibrary{ap}
	jb, _ := json.Marshal(al)
	fmt.Println(string(jb))
}

func TestLoadAircraftLibrary(t *testing.T) {
	file, err := os.Open("data/aircrafts.json")
	if err != nil {
		t.FailNow()
	}
	bytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
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

func TestAircraftPilots(t *testing.T) {
	InitGame()
	ac := NewAircraft("F14", "Default", countrycodes.UK, nato.OF2)
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", countrycodes.Germany, nato.OF2)
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", countrycodes.USA, nato.OF2)
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", countrycodes.Russia, nato.OF2)
	fmt.Println(ac)
	fmt.Println(TheRoster)
}
