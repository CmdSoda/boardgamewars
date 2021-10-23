package bgw

import (
	"fmt"
)

type DogfightPosition uint

const (
	DogfightPositionTossup DogfightPosition = 0
	DogfightPositionEnemyAtMySix = 1
	DogfightPositionBehindEnemiesTail = 2
)

func GroupDogfight(side1 []Aircraft, side2 []Aircraft) {

}

func Dogfight(aircraft1 *Aircraft, aircraft2 *Aircraft) {

	ap1 := aircraft1.GetParameters()
	ap2 := aircraft2.GetParameters()

	// In Position setzen
	// Flugzeuge mit grösseren Dogfighting-Rating haben höhere Chance.
	// 1) Kampf um die Position => Endet in einer Position
	dfa1Pos := DogfightPositionTossup
	dfa2Pos := DogfightPositionTossup
	dfr1 := Roll1D10() + int(ap1.Dogfighting)
	dfr2 := Roll1D10() + int(ap2.Dogfighting)
	dfdelta := dfr1 - dfr2
	if dfdelta > 0 {
		if dfdelta >= 3 {
			dfa1Pos = DogfightPositionBehindEnemiesTail
			dfa2Pos = DogfightPositionEnemyAtMySix
		}
	} else {
		if -dfdelta >= 3 {
			dfa1Pos = DogfightPositionEnemyAtMySix
			dfa2Pos = DogfightPositionBehindEnemiesTail
		}
	}

	fmt.Println(dfa1Pos)
	fmt.Println(dfa2Pos)

	// SRMs (Short-Range-Missles) gegeneinander einsetzen
	// 2) Abschuss der SRM

	// Gun einsetzen
	// Falls keine SRM => Einsatz der Gun

}
