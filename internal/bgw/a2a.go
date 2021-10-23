package bgw

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Air2AirWeaponParameters struct {
	EquipmentId
	Name            string
	Dogfighting     Rating // Wie gut verhält sich die Waffe im Dogfight
	BVR             Rating // Wie gut verhält sich die Waffe im BVR
	Speed           Rating // Wie schnell ist die Waffe
	Range           Rating // Wie weit fliegt die Waffe
	OrdenanceWeight Rating
}

type Air2AirWeaponLibrary []Air2AirWeaponParameters
var Air2AirLib Air2AirWeaponLibrary

func LoadAir2AirWeapons() (*Air2AirWeaponLibrary, error) {
	var err error
	file, err := os.Open("data/a2a.json")
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	a2al := Air2AirWeaponLibrary{}
	err = json.Unmarshal(bytes, &a2al)
	if err != nil {
		return nil, err
	}
	Air2AirLib = a2al
	return &a2al, nil
}

func GetAir2AirWeaponParametersFromName(name string) *Air2AirWeaponParameters {
	for _, parameters := range Air2AirLib {
		if parameters.Name == name {
			return &parameters
		}
	}
	return nil
}