package randomizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	for i := 0; i < 100; i++ {
		r := Roll1DN(1)
		assert.Equal(t, 1, r)
	}

	found1 := false
	found2 := false
	found3 := false
	for i := 0; i < 100; i++ {
		r := Roll1DN(2)
		if r == 1 {
			found1 = true
		}
		if r == 2 {
			found2 = true
		}
		if r == 3 {
			found3 = true
		}
	}
	assert.Equal(t, true, found1)
	assert.Equal(t, true, found2)
	assert.Equal(t, false, found3)
}

func TestRollRange(t *testing.T) {
	Init()
	var rolled20 bool
	var rolled21 bool
	var rolled22 bool
	var rolled23 bool
	var rolled24 bool
	var rolled25 bool
	for i := 0; i < 1000; i++ {
		r := Roll(21, 24)
		if r == 20 {
			rolled20 = true
		}
		if r == 21 {
			rolled21 = true
		}
		if r == 22 {
			rolled22 = true
		}
		if r == 23 {
			rolled23 = true
		}
		if r == 24 {
			rolled24 = true
		}
		if r == 25 {
			rolled25 = true
		}
	}
	assert.True(t, !rolled20 && rolled21 && rolled22 && rolled23 && rolled24 && !rolled25)
}
