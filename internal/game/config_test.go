package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	c, err := LoadConfig("config.json")
	assert.Nil(t, err)
	assert.Equal(t, "data", c.DataPath)
}
