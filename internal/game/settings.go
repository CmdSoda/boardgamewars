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
}

func LoadSettings(filename string) (Settings, error) {
	s := Settings{}
	file, err := os.Open(filename)
	if err != nil {
		Log.Errorf("%s not found\n", filename)
		return s, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		Log.Errorf("%s error while reading\n", filename)
		return s, err
	}
	err = json.Unmarshal(bytes, &s)
	if err != nil {
		Log.Errorf("%s error while unmarshaling\n", filename)
		return s, err
	}
	s.RandomDamage = []DamageType{}
	for _, damageString := range s.RandomDamageStrings {
		s.RandomDamage = append(s.RandomDamage, GetDamageTypeFromString(damageString))
	}
	return s, nil
}
