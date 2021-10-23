package bgw

import (
	"math/rand"
	"time"
)

func RandInit() {
	rand.Seed(time.Now().UnixNano())
}

func Roll1D10() int {
	return rand.Intn(10) + 1
}
