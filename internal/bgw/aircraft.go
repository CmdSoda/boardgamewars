package bgw

type AircraftType uint

const (
	F14   AircraftType = 0
	Mig23              = 1
	Mig27              = 2
	Su17               = 3
)

type Aircraft struct {
	Type               AircraftType
	Altitude           AltitudeBand // Aktuelle Höhe.
	CurrentPosition    Position
	NextTargetLocation Position // Das ist die Position, die das Flugzeug jetzt ansteuert.
}

type AircraftParameters struct {
	Type                  AircraftType
	Name                  string
	Nickname              string
	FirstFlight           Year
	Introduction          Year
	CombatSpeed           Rating
	CruiseSpeed           Rating
	CombatFuelConsumption Rating // Treibstoffverbrauch im Kampf pro Runde.
	CruiseFuelConsumption Rating // Treibstoffverbrauch beim Cruisen pro Runde.
	Fuel                  Rating
	MaxAltitude           AltitudeBand
	Dogfighting           Rating
	Configurations        []WeaponSystemConfiguration
	MaintenanceTime       Rating
}

type AircraftLibrary []AircraftParameters
