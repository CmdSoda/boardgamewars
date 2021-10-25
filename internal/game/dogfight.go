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
	Fighter1Position DogfightPosition
	WeaponUsed       *WeaponSystem
	Hit              bool
	DamageDone       []DamageType
}

type DogfightParameters struct {
	Aircraft
	DogfightResult
	LastDogfightPosition DogfightPosition
}

func (d DogfightResult) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "Dogfight Result: ")
	fmt.Fprintf(&sb, "%s, ", d.Fighter1Position)
	if d.WeaponUsed != nil {
		fmt.Fprintf(&sb, "WeaponUsed: %s, ", d.WeaponUsed.Name)
	} else {
		fmt.Fprint(&sb, "WeaponUsed: none, ")
	}
	fmt.Fprint(&sb, "Damage2: ")
	for _, dt := range d.DamageDone {
		fmt.Fprintf(&sb, "%s, ", dt.String())
	}
	if len(d.DamageDone) == 0 {
		fmt.Fprint(&sb, "none")
	}
	fmt.Fprint(&sb, "\n")
	return sb.String()
}
