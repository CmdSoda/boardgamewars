package randomizer

import (
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

// Roll1D10 würfel einen 10-seitigen Würfel.
func Roll1D10() int {
	return rand.Intn(10) + 1
}

// Roll1DN würfelt einen N-seitigen Würfel.
func Roll1DN(n int) int {
	return rand.Intn(n) + 1
}
