package game

import (
	"fmt"
	"testing"
)

func TestNewWorld(t *testing.T) {
	w := NewWorld()
	fmt.Println(w)
}
