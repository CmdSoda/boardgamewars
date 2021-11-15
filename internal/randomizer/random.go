package randomizer

import (
	"math/rand"
	"time"
)

func Init(seed int64) {
	if seed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(seed)
	}
}

// Roll1D10 würfel einen 10-seitigen Würfel.
func Roll1D10() int {
	return rand.Intn(10) + 1
}

// Roll1DN würfelt einen N-seitigen Würfel.
func Roll1DN(n int) int {
	return rand.Intn(n) + 1
}

func Roll(n int, m int) int {
	rng := m - n
	roll := Roll1DN(rng+1) + (n - 1)
	return roll
}
