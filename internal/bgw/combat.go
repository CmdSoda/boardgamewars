package bgw

import "fmt"

type DogfightPosition int

const (
	DogfightPositionBehindEnemiesTailOptimal DogfightPosition = 2
	DogfightPositionBehindEnemiesTail        DogfightPosition = 1
	DogfightPositionTossup                   DogfightPosition = 0
	DogfightPositionEnemyAtMySix             DogfightPosition = -1
	DogfightPositionEnemyAtMySixOptimal      DogfightPosition = -2
)

func GroupDogfight(side1 []Aircraft, side2 []Aircraft) {

}

func DogfightPerformance(rating1 Rating, rating2 Rating) DogfightPosition {
	dfr1 := Roll1D10() + int(rating1)
	dfr2 := Roll1D10() + int(rating2)
	dfdelta := dfr1 - dfr2

	if dfdelta > 0 {
		if dfdelta >= 7 {
			return DogfightPositionBehindEnemiesTailOptimal
		} else if dfdelta >= 3 {
			return DogfightPositionBehindEnemiesTail
		}
	} else {
		if -dfdelta >= 7 {
			return DogfightPositionEnemyAtMySixOptimal
		}
		if -dfdelta >= 3 {
			return DogfightPositionEnemyAtMySix
		}
	}
	return DogfightPositionTossup
}

func Dogfight(aircraft1 *Aircraft, aircraft2 *Aircraft) {
	ap1 := aircraft1.GetParameters()
	ap2 := aircraft2.GetParameters()

	// In Position setzen
	// Flugzeuge mit grösseren Dogfighting-Rating haben höhere Chance.
	// 1) Kampf um die Position => Endet in einer Position
	dfa1Pos := DogfightPerformance(ap1.Dogfighting, ap2.Dogfighting)
	fmt.Printf("aircraft1 is in %d\n", dfa1Pos)

	// SRMs (Short-Range-Missles) gegeneinander einsetzen
	// 2) Abschuss der SRM
	// Falls keine SRM => Einsatz der Gun
	if dfa1Pos > 0 {
		bestws := aircraft1.GetBestDogfightingWeapon()
		if bestws != nil {
			aircraft1.DepleteWeapon(*bestws)
			if bestws.Hit(*aircraft2, dfa1Pos) {
				dt := aircraft2.DoDamageWith(*bestws)
				fmt.Printf("aircraft2 damaged: %d\n", dt)
			} else {
				fmt.Println("aircraft1 misses")
			}
		}
	} else if -dfa1Pos > 0 {
		bestws := aircraft2.GetBestDogfightingWeapon()
		if bestws != nil {
			aircraft2.DepleteWeapon(*bestws)
			if bestws.Hit(*aircraft1, -dfa1Pos) {
				dt := aircraft1.DoDamageWith(*bestws)
				fmt.Printf("aircraft1 damaged: %d\n", dt)
			} else {
				fmt.Println("aircraft2 misses")
			}
		}
	}
}
