package bgw

type WeaponSystemType uint

const (
	FuelTank WeaponSystemType = 0
	Aim54                     = 1
	Aim9                      = 2
	Aim7                      = 3
	Gun                       = 4
)

type WeaponSystem struct {
	Name            string
	Type            WeaponSystemType
	OrdenanceWeight Rating
}

type WeaponSystemConfiguration []WeaponSystemType
