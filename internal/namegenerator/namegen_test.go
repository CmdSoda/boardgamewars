package namegenerator

import (
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSurname(t *testing.T) {
	randomizer.Init()
	sn := CreateSurname(English)
	assert.NotEqual(t, InvalidName, sn)
}
