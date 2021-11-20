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

type WinAircraftParameters struct {
	AircraftParametersId
	ConfigName string
	PilotRank  nato.Code
}

type WinStatistics struct {
	AC1Params WinAircraftParameters
	AC2Params WinAircraftParameters
	AC1Stats  WinAircraftStats
	AC2Stats  WinAircraftStats
}

type WinAircraftStats struct {
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

// AircraftVersusPerformanceList speichert Win-Statistiken:
// m[acF14]["Default"][acMig29]["Default"] = WinStatistics{ Won: 23, Lost: 67, Draw: 6 }
type AircraftVersusPerformanceList []*WinStatistics

func (w *AircraftVersusPerformanceList) Position(acid1 AircraftId, acid2 AircraftId, pos DogfightPosition) {
	ac1 := Globals.AllAircrafts[acid1]
	ac2 := Globals.AllAircrafts[acid2]
	wap1 := WinAircraftParameters{
		AircraftParametersId: ac1.AircraftParametersId,
		ConfigName:           ac1.WeaponsConfigName,
		PilotRank:            ac1.GetHighestPilotRank(),
	}
	wap2 := WinAircraftParameters{
		AircraftParametersId: ac2.AircraftParametersId,
		ConfigName:           ac2.WeaponsConfigName,
		PilotRank:            ac2.GetHighestPilotRank(),
	}
	// Die Möglichkeit besteht, dass es den Eintrag schon gibt. Suchen...
	for i, _ := range *w {
		if (*w)[i].AC1Params == wap1 && (*w)[i].AC2Params == wap2 {
			switch pos {
			case DogfightPositionAdventage:
				(*w)[i].AC1Stats.AdvantageCount = (*w)[i].AC1Stats.AdvantageCount + 1
				(*w)[i].AC2Stats.DisadvantageCount = (*w)[i].AC2Stats.DisadvantageCount + 1
			case DogfightPositionBehindEnemiesTail:
				(*w)[i].AC1Stats.BehindEnemyCount = (*w)[i].AC1Stats.BehindEnemyCount + 1
				(*w)[i].AC2Stats.EnemyAtMySixCount = (*w)[i].AC2Stats.EnemyAtMySixCount + 1
			case DogfightPositionBehindEnemiesTailOptimal:
				(*w)[i].AC1Stats.BehindEnemyOptimalCount = (*w)[i].AC1Stats.BehindEnemyOptimalCount + 1
				(*w)[i].AC2Stats.EnemyAtMySixOptimalCount = (*w)[i].AC2Stats.EnemyAtMySixOptimalCount + 1
			}
			return
		} else if (*w)[i].AC1Params == wap2 && (*w)[i].AC2Params == wap1 {
			switch pos {
			case DogfightPositionAdventage:
				(*w)[i].AC2Stats.AdvantageCount = (*w)[i].AC2Stats.AdvantageCount + 1
				(*w)[i].AC1Stats.DisadvantageCount = (*w)[i].AC1Stats.DisadvantageCount + 1
			case DogfightPositionBehindEnemiesTail:
				(*w)[i].AC2Stats.BehindEnemyCount = (*w)[i].AC2Stats.BehindEnemyCount + 1
				(*w)[i].AC1Stats.EnemyAtMySixCount = (*w)[i].AC1Stats.EnemyAtMySixCount + 1
			case DogfightPositionBehindEnemiesTailOptimal:
				(*w)[i].AC2Stats.BehindEnemyOptimalCount = (*w)[i].AC2Stats.BehindEnemyOptimalCount + 1
				(*w)[i].AC1Stats.EnemyAtMySixOptimalCount = (*w)[i].AC1Stats.EnemyAtMySixOptimalCount + 1
			}
			return
		}
	}
}

func (w *AircraftVersusPerformanceList) Win(acid1 AircraftId, acid2 AircraftId, wtype WinType) {
	ac1 := Globals.AllAircrafts[acid1]
	ac2 := Globals.AllAircrafts[acid2]
	wap1 := WinAircraftParameters{
		AircraftParametersId: ac1.AircraftParametersId,
		ConfigName:           ac1.WeaponsConfigName,
		PilotRank:            ac1.GetHighestPilotRank(),
	}
	wap2 := WinAircraftParameters{
		AircraftParametersId: ac2.AircraftParametersId,
		ConfigName:           ac2.WeaponsConfigName,
		PilotRank:            ac2.GetHighestPilotRank(),
	}
	// Die Möglichkeit besteht, dass es den Eintrag schon gibt. Suchen...
	for i, _ := range *w {
		if (*w)[i].AC1Params == wap1 && (*w)[i].AC2Params == wap2 {
			switch wtype {
			case WinTypeDraw:
				(*w)[i].AC1Stats.Draw = (*w)[i].AC1Stats.Draw + 1
				(*w)[i].AC2Stats.Draw = (*w)[i].AC2Stats.Draw + 1
			case WinTypeWon:
				(*w)[i].AC1Stats.Won = (*w)[i].AC1Stats.Won + 1
				(*w)[i].AC2Stats.Lost = (*w)[i].AC2Stats.Lost + 1
			}
			return
		} else if (*w)[i].AC1Params == wap2 && (*w)[i].AC2Params == wap1 {
			switch wtype {
			case WinTypeDraw:
				(*w)[i].AC1Stats.Draw = (*w)[i].AC1Stats.Draw + 1
				(*w)[i].AC2Stats.Draw = (*w)[i].AC2Stats.Draw + 1
			case WinTypeWon:
				(*w)[i].AC2Stats.Won = (*w)[i].AC2Stats.Won + 1
				(*w)[i].AC1Stats.Lost = (*w)[i].AC1Stats.Lost + 1
			}
			return
		}
	}

	ws := WinStatistics{
		AC1Params: wap1,
		AC2Params: wap2,
		AC1Stats: WinAircraftStats{
			Won:                     0,
			Draw:                    0,
			Lost:                    0,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		},
		AC2Stats: WinAircraftStats{
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
		ws.AC1Stats = WinAircraftStats{
			Won:                     0,
			Draw:                    1,
			Lost:                    0,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		}
		ws.AC2Stats = WinAircraftStats{
			Won:                     0,
			Draw:                    1,
			Lost:                    0,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		}
	case WinTypeWon:
		ws.AC1Stats = WinAircraftStats{
			Won:                     1,
			Draw:                    0,
			Lost:                    0,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		}
		ws.AC2Stats = WinAircraftStats{
			Won:                     0,
			Draw:                    0,
			Lost:                    1,
			AdvantageCount:          0,
			BehindEnemyCount:        0,
			BehindEnemyOptimalCount: 0,
		}
	}

	// Diese Kombination gibt es noch nicht und es muss eine erstellt werden.
	*w = append(*w, &ws)
}

func (w *AircraftVersusPerformanceList) Dump() {
	fmt.Println("Statistics: Aircraft vs Aircraft")
	for _, wstat := range *w {
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

