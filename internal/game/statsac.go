package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
)

type WinType int

const (
	WinTypeDraw WinType = 0
	WinTypeWon  WinType = 1
)

type AircraftPerformanceParameters struct {
	AircraftParametersId
	ConfigName string
	PilotRank  nato.Code
}

type AircraftVersusStatistics struct {
	AC1Params AircraftPerformanceParameters
	AC2Params AircraftPerformanceParameters
	AC1Stats  AircraftStatistics
	AC2Stats  AircraftStatistics
}

type AircraftStatistics struct {
	Won                      int
	Draw                     int
	Lost                     int
	AdvantageCount           int
	BehindEnemyCount         int
	BehindEnemyOptimalCount  int
	DisadvantageCount        int
	EnemyAtMySixCount        int
	EnemyAtMySixOptimalCount int
}

// AircraftVersusStatisticsList speichert Win-Statistiken:
// m[acF14]["Default"][acMig29]["Default"] = AircraftVersusStatistics{ Won: 23, Lost: 67, Draw: 6 }
type AircraftVersusStatisticsList []*AircraftVersusStatistics

func (avsl *AircraftVersusStatisticsList) Position(acid1 AircraftId, acid2 AircraftId, pos DogfightPosition) {
	ac1 := Globals.AllAircrafts[acid1]
	ac2 := Globals.AllAircrafts[acid2]
	wap1 := AircraftPerformanceParameters{
		AircraftParametersId: ac1.AircraftParametersId,
		ConfigName:           ac1.WeaponsConfigName,
		PilotRank:            ac1.GetHighestPilotRank(),
	}
	wap2 := AircraftPerformanceParameters{
		AircraftParametersId: ac2.AircraftParametersId,
		ConfigName:           ac2.WeaponsConfigName,
		PilotRank:            ac2.GetHighestPilotRank(),
	}
	// Die Möglichkeit besteht, dass es den Eintrag schon gibt. Suchen...
	for i := range *avsl {
		if (*avsl)[i].AC1Params == wap1 && (*avsl)[i].AC2Params == wap2 {
			switch pos {
			case DogfightPositionAdventage:
				(*avsl)[i].AC1Stats.AdvantageCount = (*avsl)[i].AC1Stats.AdvantageCount + 1
				(*avsl)[i].AC2Stats.DisadvantageCount = (*avsl)[i].AC2Stats.DisadvantageCount + 1
			case DogfightPositionBehindEnemiesTail:
				(*avsl)[i].AC1Stats.BehindEnemyCount = (*avsl)[i].AC1Stats.BehindEnemyCount + 1
				(*avsl)[i].AC2Stats.EnemyAtMySixCount = (*avsl)[i].AC2Stats.EnemyAtMySixCount + 1
			case DogfightPositionBehindEnemiesTailOptimal:
				(*avsl)[i].AC1Stats.BehindEnemyOptimalCount = (*avsl)[i].AC1Stats.BehindEnemyOptimalCount + 1
				(*avsl)[i].AC2Stats.EnemyAtMySixOptimalCount = (*avsl)[i].AC2Stats.EnemyAtMySixOptimalCount + 1
			}
			return
		} else if (*avsl)[i].AC1Params == wap2 && (*avsl)[i].AC2Params == wap1 {
			switch pos {
			case DogfightPositionAdventage:
				(*avsl)[i].AC2Stats.AdvantageCount = (*avsl)[i].AC2Stats.AdvantageCount + 1
				(*avsl)[i].AC1Stats.DisadvantageCount = (*avsl)[i].AC1Stats.DisadvantageCount + 1
			case DogfightPositionBehindEnemiesTail:
				(*avsl)[i].AC2Stats.BehindEnemyCount = (*avsl)[i].AC2Stats.BehindEnemyCount + 1
				(*avsl)[i].AC1Stats.EnemyAtMySixCount = (*avsl)[i].AC1Stats.EnemyAtMySixCount + 1
			case DogfightPositionBehindEnemiesTailOptimal:
				(*avsl)[i].AC2Stats.BehindEnemyOptimalCount = (*avsl)[i].AC2Stats.BehindEnemyOptimalCount + 1
				(*avsl)[i].AC1Stats.EnemyAtMySixOptimalCount = (*avsl)[i].AC1Stats.EnemyAtMySixOptimalCount + 1
			}
			return
		}
	}
}

func (avsl *AircraftVersusStatisticsList) Win(acid1 AircraftId, acid2 AircraftId, wtype WinType) {
	ac1 := Globals.AllAircrafts[acid1]
	ac2 := Globals.AllAircrafts[acid2]
	wap1 := AircraftPerformanceParameters{
		AircraftParametersId: ac1.AircraftParametersId,
		ConfigName:           ac1.WeaponsConfigName,
		PilotRank:            ac1.GetHighestPilotRank(),
	}
	wap2 := AircraftPerformanceParameters{
		AircraftParametersId: ac2.AircraftParametersId,
		ConfigName:           ac2.WeaponsConfigName,
		PilotRank:            ac2.GetHighestPilotRank(),
	}
	// Die Möglichkeit besteht, dass es den Eintrag schon gibt. Suchen...
	for i := range *avsl {
		if (*avsl)[i].AC1Params == wap1 && (*avsl)[i].AC2Params == wap2 {
			switch wtype {
			case WinTypeDraw:
				(*avsl)[i].AC1Stats.Draw = (*avsl)[i].AC1Stats.Draw + 1
				(*avsl)[i].AC2Stats.Draw = (*avsl)[i].AC2Stats.Draw + 1
			case WinTypeWon:
				(*avsl)[i].AC1Stats.Won = (*avsl)[i].AC1Stats.Won + 1
				(*avsl)[i].AC2Stats.Lost = (*avsl)[i].AC2Stats.Lost + 1
			}
			return
		} else if (*avsl)[i].AC1Params == wap2 && (*avsl)[i].AC2Params == wap1 {
			switch wtype {
			case WinTypeDraw:
				(*avsl)[i].AC1Stats.Draw = (*avsl)[i].AC1Stats.Draw + 1
				(*avsl)[i].AC2Stats.Draw = (*avsl)[i].AC2Stats.Draw + 1
			case WinTypeWon:
				(*avsl)[i].AC2Stats.Won = (*avsl)[i].AC2Stats.Won + 1
				(*avsl)[i].AC1Stats.Lost = (*avsl)[i].AC1Stats.Lost + 1
			}
			return
		}
	}

	ws := AircraftVersusStatistics{
		AC1Params: wap1,
		AC2Params: wap2,
		AC1Stats: AircraftStatistics{
			Won:                     0,
			Draw:                    0,
			Lost:                    0,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		},
		AC2Stats: AircraftStatistics{
			Won:                     0,
			Draw:                    0,
			Lost:                    0,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		},
	}

	switch wtype {
	case WinTypeDraw:
		ws.AC1Stats = AircraftStatistics{
			Won:                     0,
			Draw:                    1,
			Lost:                    0,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		}
		ws.AC2Stats = AircraftStatistics{
			Won:                     0,
			Draw:                    1,
			Lost:                    0,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		}
	case WinTypeWon:
		ws.AC1Stats = AircraftStatistics{
			Won:                     1,
			Draw:                    0,
			Lost:                    0,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		}
		ws.AC2Stats = AircraftStatistics{
			Won:                     0,
			Draw:                    0,
			Lost:                    1,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		}
	}

	// Diese Kombination gibt es noch nicht und es muss eine erstellt werden.
	*avsl = append(*avsl, &ws)
}

func (avsl *AircraftVersusStatisticsList) Dump() {
	fmt.Println("Statistics: Aircraft vs Aircraft")
	for _, wstat := range *avsl {
		acp1 := Globals.AllAircraftParameters[wstat.AC1Params.AircraftParametersId]
		acp2 := Globals.AllAircraftParameters[wstat.AC2Params.AircraftParametersId]
		var acwr1 float32
		var acwr2 float32
		if wstat.AC1Stats.Won+wstat.AC1Stats.Won == 0 {
			acwr1 = 0
		} else {
			acwr1 = float32(wstat.AC1Stats.Won) / float32(wstat.AC1Stats.Won+wstat.AC2Stats.Won) * 100
		}
		if wstat.AC1Stats.Won+wstat.AC2Stats.Won == 0 {
			acwr2 = 0
		} else {
			acwr2 = float32(wstat.AC2Stats.Won) / float32(wstat.AC1Stats.Won+wstat.AC2Stats.Won) * 100
		}
		samples := wstat.AC1Stats.Won + wstat.AC1Stats.Lost + wstat.AC1Stats.Draw
		fmt.Printf("%s(%.1f%%) vs %s(%.1f%%) (%d samples)\n", acp1.Name, acwr1, acp2.Name, acwr2, samples)
	}
}
