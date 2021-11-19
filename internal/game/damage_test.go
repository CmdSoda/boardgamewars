package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRollRandomDamage(t *testing.T) {
	dtl := RollRandomDamage(500, 500)
	assert.Equal(t, 3, len(dtl))
	dtl2 := RollRandomDamage(300, 500)
	assert.Equal(t, 2, len(dtl2))
	dtl3 := RollRandomDamage(100, 500)
	assert.Equal(t, 1, len(dtl3))
}

func TestRollRandomDamage2(t *testing.T) {

}
