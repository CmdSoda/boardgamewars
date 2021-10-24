package game

type FullFlight struct {
	LeadElement   FlightElement
	SecondElement FlightElement
}

type FlightElement struct {
	E1 Aircraft
	E2 Aircraft
}
