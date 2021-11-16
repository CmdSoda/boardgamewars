package game

import (
	"encoding/json"
	"github.com/CmdSoda/boardgamewars/internal/logging"
	"io/ioutil"
	"os"
)

type Config struct {
	DataPath string
}

func LoadConfig(filename string) (Config, error) {
	c := Config{}
	file, err := os.Open(filename)
	if err != nil {
		logging.Log.Errorf("%s not found\n", filename)
		return c, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		logging.Log.Errorf("%s error while reading\n", filename)
		return c, err
	}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		logging.Log.Errorf("%s error while unmarshaling\n", filename)
		return c, err
	}
	return c, nil
}
