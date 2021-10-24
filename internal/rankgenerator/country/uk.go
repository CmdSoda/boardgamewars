package country

import "github.com/CmdSoda/boardgamewars/internal/rankgenerator"

var PilotRanksUK = []rankgenerator.MilitaryFlightRank{
	{
		Name:           "Pilot Officer",
		Short:          "Plt Off",
		NatoFlightCode: rankgenerator.NatoFlightCodeOF1,
	},
	{
		Name:           "Flying Officer",
		Short:          "Fg Off",
		NatoFlightCode: rankgenerator.NatoFlightCodeOF1,
	},
	{
		Name:           "Flight Lieutenant",
		Short:          "Flt Lt",
		NatoFlightCode: rankgenerator.NatoFlightCodeOF2,
	},
	{
		Name:           "Squadron Leader",
		Short:          "Sqn Ldr",
		NatoFlightCode: rankgenerator.NatoFlightCodeOF3,
	},
	{
		Name:           "Wing Commander",
		Short:          "Wg Cdr",
		NatoFlightCode: rankgenerator.NatoFlightCodeOF4,
	},
	{
		Name:           "Group Captain",
		Short:          "Gp Cpt",
		NatoFlightCode: rankgenerator.NatoFlightCodeOF5,
	},
}
