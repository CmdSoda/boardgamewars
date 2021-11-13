package game

import (
	"fmt"
	"strings"
)

type DogfightPosition int

const (
	DogfightPositionBehindEnemiesTailOptimal DogfightPosition = 3
	DogfightPositionBehindEnemiesTail        DogfightPosition = 2
	DogfightPositionAdventage                DogfightPosition = 1
	DogfightPositionTossup                   DogfightPosition = 0
	DogfightPositionIgnore                   DogfightPosition = 0
	DogfightPositionDisadvantage             DogfightPosition = -1
	DogfightPositionEnemyAtMySix             DogfightPosition = -2
	DogfightPositionEnemyAtMySixOptimal      DogfightPosition = -3
)

func (dp DogfightPosition) String() string {
	switch dp {
	case DogfightPositionBehindEnemiesTailOptimal:
		return "BehindEnemyOptimal"
	case DogfightPositionBehindEnemiesTail:
		return "BehindEnemy"
	case DogfightPositionAdventage:
		return "Adventage"
	case DogfightPositionTossup:
		return "Tossup"
	case DogfightPositionDisadvantage:
		return "Disadventage"
	case DogfightPositionEnemyAtMySix:
		return "EnemyAtMySix"
	case DogfightPositionEnemyAtMySixOptimal:
		return "EnemyAtMySixOptimal"
	}
	return "Invalid"
}

type DogfightResult struct {
	Round            int
	Position         DogfightPosition
	WeaponUsed       *WeaponSystem
	Hit              bool
	DamageConflicted []DamageType
}

// DogfightSetup wird initial für einen Kampf benötigt. Aus diesem struct entsteht dann Dogfight.
type DogfightSetup struct {
	TeamBlue AircraftIdList
	TeamRed  AircraftIdList
}

// DogfightGroup wird aus dem struct Dogfight erstellt. Je mehr Flugzeuge in der Dogfight Warteliste sind, desto mehr
// DogfightGroup-Objekte werden erzeugt.
type DogfightGroup struct {
	BlueFighter *AircraftId
	BlueSupport *AircraftId // optional
	RedFighter  *AircraftId
	RedSupport  *AircraftId // optional
}

// Dogfight wird aus einem DogfightSetup initialisiert. Während des Kampfes werden so viele DogfightGroup erstellt, wie
// es möglich ist.
type Dogfight struct {
	Groups          []DogfightGroup
	TeamBlueWaiting AircraftIdList
	TeamRedWaiting  AircraftIdList
}

//goland:noinspection ALL
func (d DogfightResult) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d.ExecuteDogfight Result: ", d.Round)
	fmt.Fprintf(&sb, "%s", d.Position)
	if d.WeaponUsed != nil {
		fmt.Fprintf(&sb, ", HitWith: %s", d.WeaponUsed.WeaponSystemName)
		if len(d.DamageConflicted) > 0 {
			fmt.Fprint(&sb, ", DamageConflicted: ")
			for _, dt := range d.DamageConflicted {
				fmt.Fprintf(&sb, "%s, ", dt.String())
			}
		} else {
			fmt.Fprint(&sb, ", no damage")
		}
	}
	fmt.Fprint(&sb, "\n")
	return sb.String()
}
