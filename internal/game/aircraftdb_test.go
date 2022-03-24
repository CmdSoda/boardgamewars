package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAndDropTable(t *testing.T) {
	assert.Nil(t, InitGame(0))
	assert.Nil(t, DropAircraftTable())
	assert.Nil(t, CreateAircraftTable())
	CloseDatabase()
}
