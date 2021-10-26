package game

import (
	"encoding/json"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)

type Air2AirWeaponParameters struct {
	Id              uuid.UUID
	Name            string
	Dogfighting     Rating // Wie gut verhält sich die Waffe im Dogfight
	BVR             Rating // Wie gut verhält sich die Waffe im BVR
	Speed           Rating // Wie schnell ist die Waffe
	Range           Rating // Wie weit fliegt die Waffe
	Damage          Hitpoints
	OrdenanceWeight Rating
	Tags            []string
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

func (awp Air2AirWeaponParameters) Hit(target Aircraft, dfp DogfightPosition) bool {
	wep := awp.Dogfighting
	if dfp == DogfightPositionBehindEnemiesTailOptimal {
		wep += 2
	}
	dfpw := DogfightPerformance(wep, DogfightPositionIgnore, target.GetParameters().Dogfighting, DogfightPositionIgnore)
	return dfpw > 0
}

func (awp Air2AirWeaponParameters) DoRandomDamage() Hitpoints {
	dr := randomizer.Roll1D10()
	if dr >= 9 {
		return awp.Damage
	} else if dr >= 7 {
		return Hitpoints(float32(awp.Damage) * 0.50)
	} else if dr >= 5 {
		return Hitpoints(float32(awp.Damage) * 0.25)
	} else if dr >= 3 {
		return Hitpoints(float32(awp.Damage) * 0.10)
	} else {
		return 0
	}
}
