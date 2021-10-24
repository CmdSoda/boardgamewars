package namegenerator

import "github.com/CmdSoda/boardgamewars/internal/randomizer"

type CountryCode int

const (
	English     CountryCode = 0
	German      CountryCode = 1
	InvalidName string      = "InvalidName"
)

func CreateSurname(cc CountryCode) string {
	switch cc {
	case English:
		r := randomizer.Roll1DN(len(englishSurnames))
		return englishSurnames[r-1]
	case German:
		r := randomizer.Roll1DN(len(germanSurnames))
		return germanSurnames[r-1]
	}
	return InvalidName
}
