package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ScenarioSettings struct {
	// Durchmesser eines Hex in Kilometern.
	HexDiameterKm float32
}

func LoadScenarioSettings(filename string) (ScenarioSettings, error) {
	s := ScenarioSettings{}
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
	return s, nil
}

//goland:noinspection GoUnhandledErrorResult
func (s ScenarioSettings) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "HexDiameterKm = %f\n", s.HexDiameterKm)
	return b.String()
}
