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
	acid1 AircraftId, ldp1 DogfightPosition,
	acid2 AircraftId, ldp2 DogfightPosition) (DogfightResult, DogfightResult) {
	var dfr1 DogfightResult
	var dfr2 DogfightResult
	ac1 := Globals.AllAircrafts[acid1]
	ac2 := Globals.AllAircrafts[acid2]
	ap1 := ac1.GetParameters()
	ap2 := ac2.GetParameters()

	// In FloatPosition setzen
	// Flugzeuge mit grösseren Dogfighting-Rating haben höhere Chance.
	// 1) Kampf um die FloatPosition => Endet in einer FloatPosition
	dfa1Pos := DogfightPerformance(ap1.Dogfighting, ldp1, ap2.Dogfighting, ldp2)
	dfr1.Position = dfa1Pos
	dfr2.Position = -dfa1Pos

	// SRMs (Short-Range-Missles) gegeneinander einsetzen
	// 2) Abschuss der SRM
	// Falls keine SRM => Einsatz der Gun
	if dfa1Pos >= DogfightPositionBehindEnemiesTail {
		bestws, exist := ac1.GetBestDogfightingWeapon()
		dfr1.WeaponUsed = &bestws
		if exist {
			ac1.DepleteWeapon(bestws)
			if bestws.Hit(acid2, dfa1Pos) {
				dfr1.Hit = true
				dt := ac2.DoDamageWith(bestws)
				dfr1.DamageConflicted = append(dfr1.DamageConflicted, dt)
			}
		}
	} else if -dfa1Pos >= DogfightPositionBehindEnemiesTail {
		bestws, exist := ac2.GetBestDogfightingWeapon()
		dfr2.WeaponUsed = &bestws
		if exist {
			ac2.DepleteWeapon(bestws)
			if bestws.Hit(acid1, -dfa1Pos) {
				dfr2.Hit = true
				dt := ac1.DoDamageWith(bestws)
				dfr2.DamageConflicted = append(dfr2.DamageConflicted, dt)
			}
		}
	}
	return dfr1, dfr2
}

func Sim10Rounds(acid1 AircraftId, acid2 AircraftId) (*[]DogfightResult, *[]DogfightResult) {
	drl1 := make([]DogfightResult, 0)
	drl2 := make([]DogfightResult, 0)

	ldp1 := DogfightPositionTossup
	ldp2 := DogfightPositionTossup

	for i := 0; i < 10; i++ {
		dr1, dr2 := ExecuteDogfight(acid1, ldp1, acid2, ldp2)
		dr1.Round = i
		dr2.Round = i
		ldp1 = dr1.Position
		ldp2 = dr2.Position
		drl1 = append(drl1, dr1)
		drl2 = append(drl2, dr2)
	}
	return &drl1, &drl2
}
