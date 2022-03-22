package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAndDropAirbases(t *testing.T) {
	assert.Nil(t, InitGame(0))
	assert.Nil(t, DropAirbaseTable())
	assert.Nil(t, CreateAirbaseTable())
	CloseDatabase()
}
