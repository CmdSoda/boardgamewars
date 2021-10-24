package country

import "github.com/CmdSoda/boardgamewars/internal/military"

var PilotRanksUK = []military.FlightRank{
	{
		Name:           "Pilot Officer",
		Short:          "Plt Off",
		NatoFlightCode: military.NatoFlightCodeOF1,
	},
	{
		Name:           "Flying Officer",
		Short:          "Fg Off",
		NatoFlightCode: military.NatoFlightCodeOF1,
	},
	{
		Name:           "Flight Lieutenant",
		Short:          "Flt Lt",
		NatoFlightCode: military.NatoFlightCodeOF2,
	},
	{
		Name:           "Squadron Leader",
		Short:          "Sqn Ldr",
		NatoFlightCode: military.NatoFlightCodeOF3,
	},
	{
		Name:           "Wing Commander",
		Short:          "Wg Cdr",
		NatoFlightCode: military.NatoFlightCodeOF4,
	},
	{
		Name:           "Group Captain",
		Short:          "Gp Cpt",
		NatoFlightCode: military.NatoFlightCodeOF5,
	},
}
