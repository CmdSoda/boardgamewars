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
	Dogfighting     Rating // Wie gut verhält sich die Waffe im ExecuteDogfight
	BVR             Rating // Wie gut verhält sich die Waffe im BVR
	Speed           Rating // Wie schnell ist die Waffe
	Range           Rating // Wie weit fliegt die Waffe
	Damage          Hitpoints
	OrdenanceWeight Rating
	Tags            []string
}

type Air2AirWeaponLibrary []Air2AirWeaponParameters

func LoadAir2AirWeapons() error {
	var err error
	file, err := os.Open("data/air2air.json")
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	a2al := Air2AirWeaponLibrary{}
	err = json.Unmarshal(bytes, &a2al)
	if err != nil {
		return err
	}
	Globals.Air2AirWeaponLibrary = a2al
	return nil
}

func GetAir2AirWeaponParametersFromName(name string) (Air2AirWeaponParameters, bool) {
	for _, parameters := range Globals.Air2AirWeaponLibrary {
		if parameters.Name == name {
			return parameters, true
		}
	}
	return Air2AirWeaponParameters{}, false
}

func (awp Air2AirWeaponParameters) Hit(target AircraftId, dfp DogfightPosition) bool {
	wep := awp.Dogfighting
	if dfp == DogfightPositionBehindEnemiesTailOptimal {
		wep += 2
	}
	targetac := Globals.AllAircrafts[target]
	dfpw := SimulateDogfightPosition(wep, DogfightPositionIgnore, targetac.GetParameters().Dogfighting, DogfightPositionIgnore)
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
