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