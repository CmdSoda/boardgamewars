package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertKmToNm(t *testing.T) {
	assert.Equal(t, 0.539957, ConvertKmToNm(1))
}
