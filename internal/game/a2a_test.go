package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateAir2AirWeaponsLibrary(t *testing.T) {
	ap := Air2AirWeaponParameters{}
	ap.EquipmentId = 0
	ap.Name = "Aim-7"
	ap.Dogfighting = 9
	ap.BVR = 0
	ap.Speed = 9
	ap.Range = 10
	ap.OrdenanceWeight = 4
	al := Air2AirWeaponLibrary{ap}
	jb, _ := json.Marshal(al)
	fmt.Println(string(jb))
}

func TestLoadAir2AirLibrary(t *testing.T) {
	var err error
	file, err := os.Open("data/a2a.json")
	if err != nil {
		t.FailNow()
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.FailNow()
	}
	a2al := Air2AirWeaponLibrary{}
	err = json.Unmarshal(bytes, &a2al)
	if err != nil {
		t.FailNow()
	}

	if len(a2al) == 0 {
		t.FailNow()
	}
}
