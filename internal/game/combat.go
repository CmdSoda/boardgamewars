package game

import (
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

func GroupDogfight(side1 []Aircraft, side2 []Aircraft) {

}

func DogfightPerformance(rating1 Rating, rating2 Rating) DogfightPosition {
	dfr1 := randomizer.Roll1D10() + int(rating1)
	dfr2 := randomizer.Roll1D10() + int(rating2)
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

func Dogfight(aircraft1 *Aircraft, aircraft2 *Aircraft) (DogfightResult, DogfightResult) {
	ap1 := aircraft1.GetParameters()
	ap2 := aircraft2.GetParameters()

	dr1 := DogfightResult{}
	dr2 := DogfightResult{}

	// In Position setzen
	// Flugzeuge mit grösseren Dogfighting-Rating haben höhere Chance.
	// 1) Kampf um die Position => Endet in einer Position
	dfa1Pos := DogfightPerformance(ap1.Dogfighting, ap2.Dogfighting)
	dr1.Fighter1Position = dfa1Pos
	dr2.Fighter1Position = -dfa1Pos

	// SRMs (Short-Range-Missles) gegeneinander einsetzen
	// 2) Abschuss der SRM
	// Falls keine SRM => Einsatz der Gun
	if dfa1Pos > 0 {
		bestws := aircraft1.GetBestDogfightingWeapon()
		dr1.WeaponUsed = bestws
		if bestws != nil {
			aircraft1.DepleteWeapon(*bestws)
			if bestws.Hit(*aircraft2, dfa1Pos) {
				dr1.Hit = true
				dt := aircraft2.DoDamageWith(*bestws)
				dr1.DamageDone = append(dr1.DamageDone, dt)
			}
		}
	} else if -dfa1Pos > 0 {
		bestws := aircraft2.GetBestDogfightingWeapon()
		dr2.WeaponUsed = bestws
		if bestws != nil {
			aircraft2.DepleteWeapon(*bestws)
			if bestws.Hit(*aircraft1, -dfa1Pos) {
				dr2.Hit = true
				dt := aircraft1.DoDamageWith(*bestws)
				dr2.DamageDone = append(dr2.DamageDone, dt)
			}
		}
	}

	return dr1, dr2
}
