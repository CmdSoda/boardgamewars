package rankgenerator

type NatoFlightCode int

const (
	NatoFlightCodeOF1 NatoFlightCode = 1
	NatoFlightCodeOF2 NatoFlightCode = 2
	NatoFlightCodeOF3 NatoFlightCode = 3
	NatoFlightCodeOF4 NatoFlightCode = 4
	NatoFlightCodeOF5 NatoFlightCode = 5
	NatoFlightCodeOF6 NatoFlightCode = 6
)

type MilitaryFlightRank struct {
	Name string
	Short string
	NatoFlightCode
}

