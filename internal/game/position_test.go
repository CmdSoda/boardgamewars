package game

import (
	"fmt"
	"testing"
)

func TestPosition_String(t *testing.T) {
	pos := FloatPosition{X: 0.54, Y: 23.7}
	fmt.Println(pos)
}
