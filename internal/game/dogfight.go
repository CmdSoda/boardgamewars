package game

import (
	"fmt"
	"strings"
)

type DogfightPosition int

const (
	DogfightPositionBehindEnemiesTailOptimal DogfightPosition = 2
	DogfightPositionBehindEnemiesTail        DogfightPosition = 1
	DogfightPositionTossup                   DogfightPosition = 0
	DogfightPositionEnemyAtMySix             DogfightPosition = -1
	DogfightPositionEnemyAtMySixOptimal      DogfightPosition = -2
)

func (dp DogfightPosition) String() string {
	switch dp {
	case DogfightPositionBehindEnemiesTailOptimal:
		return "BehindEnemyOptimal"
	case DogfightPositionBehindEnemiesTail:
		return "BehindEnemy"
	case DogfightPositionTossup:
		return "Tossup"
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

func (d DogfightResult) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "Dogfight Result:\n")
	fmt.Fprintf(&sb, "Position: %s\n", d.Fighter1Position)
	if d.WeaponUsed != nil {
		fmt.Fprintf(&sb, "WeaponUsed: %s\n", d.WeaponUsed.Name)
	} else {
		fmt.Fprint(&sb, "WeaponUsed: nil\n")
	}
	fmt.Fprint(&sb, "Damage: ")
	for _, dt := range d.DamageDone {
		fmt.Fprintf(&sb, "%s, ", dt.String())
	}
	fmt.Fprint(&sb, "\n")
	return sb.String()
}
