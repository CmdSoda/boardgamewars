package game

// FlightElement 2 aircrafts
type FlightElement struct {
	E1 Aircraft
	E2 Aircraft
}

// FullFlight 4 aircrafts
type FullFlight struct {
	LeadElement   FlightElement
	SecondElement FlightElement
}

// Squadron 16 aircrafts
type Squadron struct {
	Flight1 FullFlight
	Flight2 FullFlight
	Flight3 FullFlight
	Flight4 FullFlight
}
