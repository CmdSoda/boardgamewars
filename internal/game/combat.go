package game

import (
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

func DogfightPerformance(rating1 Rating, lastPosition1 DogfightPosition,
	rating2 Rating, lastPosition2 DogfightPosition) DogfightPosition {
	dfr1 := randomizer.Roll1D10() + int(rating1)
	dfr2 := randomizer.Roll1D10() + int(rating2)
	dfdelta := dfr1 - dfr2

	if lastPosition1 == DogfightPositionAdventage {
		dfdelta = dfdelta + 3
	}

	if lastPosition2 == DogfightPositionAdventage {
		dfdelta = dfdelta - 3
	}

	if dfdelta > 0 {
		if dfdelta >= 7 {
			return DogfightPositionBehindEnemiesTailOptimal
		} else if dfdelta >= 4 {
			return DogfightPositionBehindEnemiesTail
		} else if dfdelta >= 2 {
			return DogfightPositionAdventage
		}
	} else {
		if -dfdelta >= 7 {
			return DogfightPositionEnemyAtMySixOptimal
		} else if -dfdelta >= 4 {
			return DogfightPositionEnemyAtMySix
		} else if -dfdelta >= 2 {
			return DogfightPositionDisadvantage
		}
	}
	return DogfightPositionTossup
}

// ExecuteDogfight Eine Runde im Luftkampf. Etwa 10 Sekunden dauer.
func ExecuteDogfight(
	aircraft1 *Aircraft, ldp1 DogfightPosition,
	aircraft2 *Aircraft, ldp2 DogfightPosition) (DogfightResult, DogfightResult) {
	var dfr1 DogfightResult
	var dfr2 DogfightResult
	ap1 := aircraft1.GetParameters()
	ap2 := aircraft2.GetParameters()

	// In Position setzen
	// Flugzeuge mit grösseren Dogfighting-Rating haben höhere Chance.
	// 1) Kampf um die Position => Endet in einer Position
	dfa1Pos := DogfightPerformance(ap1.Dogfighting, ldp1, ap2.Dogfighting, ldp2)
	dfr1.Position = dfa1Pos
	dfr2.Position = -dfa1Pos

	// SRMs (Short-Range-Missles) gegeneinander einsetzen
	// 2) Abschuss der SRM
	// Falls keine SRM => Einsatz der Gun
	if dfa1Pos >= DogfightPositionBehindEnemiesTail {
		bestws := aircraft1.GetBestDogfightingWeapon()
		dfr1.WeaponUsed = bestws
		if bestws != nil {
			aircraft1.DepleteWeapon(*bestws)
			if bestws.Hit(*aircraft2, dfa1Pos) {
				dfr1.Hit = true
				dt := aircraft2.DoDamageWith(*bestws)
				dfr1.DamageConflicted = append(dfr1.DamageConflicted, dt)
			}
		}
	} else if -dfa1Pos >= DogfightPositionBehindEnemiesTail {
		bestws := aircraft2.GetBestDogfightingWeapon()
		dfr2.WeaponUsed = bestws
		if bestws != nil {
			aircraft2.DepleteWeapon(*bestws)
			if bestws.Hit(*aircraft1, -dfa1Pos) {
				dfr2.Hit = true
				dt := aircraft1.DoDamageWith(*bestws)
				dfr2.DamageConflicted = append(dfr2.DamageConflicted, dt)
			}
		}
	}
	return dfr1, dfr2
}

func Sim10Rounds(aircraft1 *Aircraft, aircraft2 *Aircraft) (*[]DogfightResult, *[]DogfightResult) {
	drl1 := make([]DogfightResult, 0)
	drl2 := make([]DogfightResult, 0)

	ldp1 := DogfightPositionTossup
	ldp2 := DogfightPositionTossup

	for i := 0; i < 10; i++ {
		dr1, dr2 := ExecuteDogfight(aircraft1, ldp1, aircraft2, ldp2)
		dr1.Round = i
		dr2.Round = i
		ldp1 = dr1.Position
		ldp2 = dr2.Position
		drl1 = append(drl1, dr1)
		drl2 = append(drl2, dr2)
	}
	return &drl1, &drl2
}
