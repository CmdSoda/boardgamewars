package game

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadScenarioSettings(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	s, err := LoadScenarioSettings(Globals.Startup.DataPath + "scenario_example1.json")
	assert.Nil(t, err)
	fmt.Println(s)
	assert.Greater(t, s.HexDiameterKm, float32(0))
}
