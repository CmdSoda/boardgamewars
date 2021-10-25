package game

import (
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

func DogfightPerformance(rating1 Rating, lastPosition1 DogfightPosition,
	rating2 Rating, lastPosition2 DogfightPosition) DogfightPosition {
	dfr1 := randomizer.Roll1D10() + int(rating1)
	dfr2 := randomizer.Roll1D10() + int(rating2)
	dfdelta := dfr1 - dfr2

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

// Dogfight Eine Runde im Luftkampf. Etwa 10 Sekunden dauer.
func Dogfight(df1 *DogfightParameters, df2 *DogfightParameters) {
	ap1 := df1.Aircraft.GetParameters()
	ap2 := df1.Aircraft.GetParameters()

	// In Position setzen
	// Flugzeuge mit grösseren Dogfighting-Rating haben höhere Chance.
	// 1) Kampf um die Position => Endet in einer Position
	dfa1Pos := DogfightPerformance(ap1.Dogfighting, df1.LastDogfightPosition, ap2.Dogfighting, df2.LastDogfightPosition)
	df1.DogfightResult.Fighter1Position = dfa1Pos
	df2.DogfightResult.Fighter1Position = -dfa1Pos

	// SRMs (Short-Range-Missles) gegeneinander einsetzen
	// 2) Abschuss der SRM
	// Falls keine SRM => Einsatz der Gun
	if dfa1Pos >= DogfightPositionBehindEnemiesTail {
		bestws := df1.Aircraft.GetBestDogfightingWeapon()
		df1.DogfightResult.WeaponUsed = bestws
		if bestws != nil {
			df1.Aircraft.DepleteWeapon(*bestws)
			if bestws.Hit(df2.Aircraft, dfa1Pos) {
				df1.DogfightResult.Hit = true
				dt := df2.Aircraft.DoDamageWith(*bestws)
				df1.DogfightResult.DamageDone = append(df1.DogfightResult.DamageDone, dt)
			}
		}
	} else if -dfa1Pos >= DogfightPositionBehindEnemiesTail {
		bestws := df2.Aircraft.GetBestDogfightingWeapon()
		df2.DogfightResult.WeaponUsed = bestws
		if bestws != nil {
			df2.Aircraft.DepleteWeapon(*bestws)
			if bestws.Hit(df1.Aircraft, -dfa1Pos) {
				df2.DogfightResult.Hit = true
				dt := df1.Aircraft.DoDamageWith(*bestws)
				df2.DogfightResult.DamageDone = append(df2.DogfightResult.DamageDone, dt)
			}
		}
	}
}

func Sim10Rounds(aircraft1 *Aircraft, aircraft2 *Aircraft) ([]DogfightResult, []DogfightResult) {
	drl1 := make([]DogfightResult, 0)
	drl2 := make([]DogfightResult, 0)

	dr1 := DogfightResult{}
	dr2 := DogfightResult{}
	dfp1 := &DogfightParameters{
		Aircraft:             *aircraft1,
		DogfightResult:       dr1,
		LastDogfightPosition: 0,
	}
	dfp2 := &DogfightParameters{
		Aircraft:             *aircraft2,
		DogfightResult:       dr2,
		LastDogfightPosition: 0,
	}

	for i := 0; i < 10; i++ {
		Dogfight(dfp1, dfp2)
		dfp1.LastDogfightPosition = dfp1.DogfightResult.Fighter1Position
		dfp2.LastDogfightPosition = dfp2.DogfightResult.Fighter1Position
		drl1 = append(drl1, dfp1.DogfightResult)
		drl2 = append(drl2, dfp2.DogfightResult)
	}
	return drl1, drl2
}
