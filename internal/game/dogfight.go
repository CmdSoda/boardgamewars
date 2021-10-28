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

type DogfightSituation struct {
	BlueFighter Aircraft
	BlueSupport *Aircraft // optional
	RedFighter  Aircraft
	RedSupport  *Aircraft // optional
}

type DogfightSetup struct {
	TeamBlue   []Aircraft
	TeamRed    []Aircraft
	Situations []DogfightSituation
}

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

type DogfightParameters struct {
	Aircraft
	DogfightResult
	LastDogfightPosition DogfightPosition
}

func (d DogfightResult) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d.Dogfight Result: ", d.Round)
	fmt.Fprintf(&sb, "%s", d.Position)
	if d.WeaponUsed != nil {
		fmt.Fprintf(&sb, ", HitWith: %s", d.WeaponUsed.Name)
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
