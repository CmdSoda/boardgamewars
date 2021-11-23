package game

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Startup struct {
	DataPath string
	LogLevel string
}

func endsWithSlash(path string) bool {
	return strings.HasSuffix(path, "/")
}

func LoadStartup(filename string) (Startup, error) {
	c := Startup{}
	file, err := os.Open(filename)
	if err != nil {
		Log.Errorf("%s not found\n", filename)
		return c, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		Log.Errorf("%s error while reading\n", filename)
		return c, err
	}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		Log.Errorf("%s error while unmarshaling\n", filename)
		return c, err
	}

	if !endsWithSlash(c.DataPath) {
		c.DataPath = c.DataPath + "/"
	}

	return c, nil
}
