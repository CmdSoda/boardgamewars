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

func LoadSettings(filename string) error {
	dataPathfilename := Globals.Startup.DataPath + filename
	file, err := os.Open(dataPathfilename)
	if err != nil {
		Log.Errorf("%s not found\n", dataPathfilename)
		return err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		Log.Errorf("%s error while reading\n", dataPathfilename)
		return err
	}
	err = json.Unmarshal(bytes, &Globals.Settings)
	if err != nil {
		Log.Errorf("%s error while unmarshaling\n", dataPathfilename)
		return err
	}
	Globals.Settings.RandomDamage = []DamageType{}
	for _, damageString := range Globals.Settings.RandomDamageStrings {
		Globals.Settings.RandomDamage = append(Globals.Settings.RandomDamage, GetDamageTypeFromString(damageString))
	}
	return nil
}
