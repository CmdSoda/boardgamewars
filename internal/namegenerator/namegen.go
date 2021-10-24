package namegenerator

import (
	"github.com/CmdSoda/boardgamewars/internal/namegenerator/country"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

type CountryCode int

const (
	UK     CountryCode = 0
	German CountryCode = 1
	InvalidParameter string      = "InvalidParameter"
)

func CreateSurname(cc CountryCode) string {
	switch cc {
	case UK:
		r := randomizer.Roll1DN(len(country.SurnamesUK))
		return country.SurnamesUK[r-1]
	case German:
		r := randomizer.Roll1DN(len(country.SurnamesGermany))
		return country.SurnamesGermany[r-1]
	}
	return InvalidParameter
}

func CreateMaleFirstname(cc CountryCode) string {
	switch cc {
	case UK:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesUK))
		return country.MaleFirstNamesUK[r-1]
	case German:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesGermany))
		return country.MaleFirstNamesGermany[r-1]
	}
	return InvalidParameter
}

func CreateFemaleFirstname(cc CountryCode) string {
	switch cc {
	case UK:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesUK))
		return country.FemaleFirstNamesUK[r-1]
	case German:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesGermany))
		return country.FemaleFirstNamesGermany[r-1]
	}
	return InvalidParameter
}

func CreateMaleFullName(cc CountryCode) string {
	switch cc {
	case UK:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesUK))
		firstname := country.MaleFirstNamesUK[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesUK))
		surname := country.SurnamesUK[r-1]
		return firstname + " " + surname
	case German:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesGermany))
		firstname := country.MaleFirstNamesGermany[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesGermany))
		surname := country.SurnamesGermany[r-1]
		return firstname + " " + surname
	}
	return InvalidParameter
}

func CreateFemaleFullName(cc CountryCode) string {
	switch cc {
	case UK:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesUK))
		firstname := country.FemaleFirstNamesUK[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesUK))
		surname := country.SurnamesUK[r-1]
		return firstname + " " + surname
	case German:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesGermany))
		firstname := country.FemaleFirstNamesGermany[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesGermany))
		surname := country.SurnamesGermany[r-1]
		return firstname + " " + surname
	}
	return InvalidParameter
}
