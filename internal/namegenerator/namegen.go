package namegenerator

import (
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator/country"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

func CreateSurname(cc countrycodes.Code) string {
	switch cc {
	case countrycodes.UK:
		r := randomizer.Roll1DN(len(country.SurnamesUK))
		return country.SurnamesUK[r-1]
	case countrycodes.Germany:
		r := randomizer.Roll1DN(len(country.SurnamesGermany))
		return country.SurnamesGermany[r-1]
	}
	return countrycodes.InvalidParameter
}

func CreateMaleFirstname(cc countrycodes.Code) string {
	switch cc {
	case countrycodes.UK:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesUK))
		return country.MaleFirstNamesUK[r-1]
	case countrycodes.Germany:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesGermany))
		return country.MaleFirstNamesGermany[r-1]
	}
	return countrycodes.InvalidParameter
}

func CreateFemaleFirstname(cc countrycodes.Code) string {
	switch cc {
	case countrycodes.UK:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesUK))
		return country.FemaleFirstNamesUK[r-1]
	case countrycodes.Germany:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesGermany))
		return country.FemaleFirstNamesGermany[r-1]
	}
	return countrycodes.InvalidParameter
}

func CreateMaleFullName(cc countrycodes.Code) string {
	switch cc {
	case countrycodes.UK:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesUK))
		firstname := country.MaleFirstNamesUK[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesUK))
		surname := country.SurnamesUK[r-1]
		return firstname + " " + surname
	case countrycodes.Germany:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesGermany))
		firstname := country.MaleFirstNamesGermany[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesGermany))
		surname := country.SurnamesGermany[r-1]
		return firstname + " " + surname
	case countrycodes.USA:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesUK))
		firstname := country.MaleFirstNamesUK[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesUK))
		surname := country.SurnamesUK[r-1]
		return firstname + " " + surname
	}
	return countrycodes.InvalidParameter
}

func CreateFemaleFullName(cc countrycodes.Code) string {
	switch cc {
	case countrycodes.UK:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesUK))
		firstname := country.FemaleFirstNamesUK[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesUK))
		surname := country.SurnamesUK[r-1]
		return firstname + " " + surname
	case countrycodes.Germany:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesGermany))
		firstname := country.FemaleFirstNamesGermany[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesGermany))
		surname := country.SurnamesGermany[r-1]
		return firstname + " " + surname
	}
	return countrycodes.InvalidParameter
}
