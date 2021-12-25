package game

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Settings struct {
	DogfightPositionAdventageBonus                    int
	DogfightPositionBehindEnemiesTailOptimalThreshold int
	DogfightPositionBehindEnemiesTailThreshold        int
	DogfightPositionAdventageThreshold                int
	RandomDamageStrings                               []string
	RandomDamage                                      []DamageType
	DamageMaintenanceMultiplier                       map[string]float64
	RepairTimePerDamageTypeBase                       float64
}

func LoadSettings(filename string) (Settings, error) {
	s := Settings{}
	dataPathfilename := Globals.Startup.DataPath + filename
	file, err := os.Open(dataPathfilename)
	if err != nil {
		Log.Errorf("%s not found\n", dataPathfilename)
		return s, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		Log.Errorf("%s error while reading\n", dataPathfilename)
		return s, err
	}
	err = json.Unmarshal(bytes, &s)
	if err != nil {
		Log.Errorf("%s error while unmarshaling\n", dataPathfilename)
		return s, err
	}
	s.RandomDamage = []DamageType{}
	for _, damageString := range s.RandomDamageStrings {
		s.RandomDamage = append(s.RandomDamage, GetDamageTypeFromString(damageString))
	}
	return s, nil
}
