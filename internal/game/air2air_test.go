package game

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateAir2AirWeaponsLibrary(t *testing.T) {
	ap := Air2AirWeaponParameters{}
	ap.Id = uuid.New()
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
	var file *os.File
	var bytes []byte

	file, err = os.Open("data/air2air.json")
	if err != nil {
		t.FailNow()
	}
	bytes, err = ioutil.ReadAll(file)
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

func TestGetAir2AirWeaponParametersFromName(t *testing.T) {
	assert.Nil(t, InitGame())
	a2awp, exist := GetAir2AirWeaponParametersFromName("ubekannt")
	assert.False(t, exist)
	assert.Equal(t, Air2AirWeaponParameters{}, a2awp)
}
