package military

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

type FlightRank struct {
	Name  string
	Short string
	NatoOfficerCode
}

func (f FlightRank) String() string {
	if f.Short != "" {
		return fmt.Sprintf("%s", f.Short)
	} else {
		return f.Name
	}
}

type FlightRankList []FlightRank

func (frl FlightRankList) RandomPick() FlightRank {
	roll := randomizer.Roll1DN(len(frl))
	return frl[roll-1]
}

func GetAllRanksWithNatoCode(list []FlightRank, noc NatoOfficerCode) FlightRankList {
	frl := make(FlightRankList, 0)
	for _, rank := range list {
		if rank.NatoOfficerCode == noc {
			frl = append(frl, rank)
		}
	}
	return frl
}

func GetFlightRank(list []FlightRank, noc NatoOfficerCode) FlightRank {
	frl := GetAllRanksWithNatoCode(list, noc)
	return frl.RandomPick()
}

func NewRank(af AirForce, noc NatoOfficerCode) FlightRank {
	switch af {
	case AirForceRAF:
		return GetFlightRank(PilotRanksRAF, noc)
	case AirForceGAF:
		return GetFlightRank(PilotRanksGAF, noc)
	case AirForceUSAF:
		return GetFlightRank(PilotRanksUSAF, noc)
	}
	return FlightRank{}
}
