package namegenerator

import (
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator/country"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"strings"
)

func CreateSurname(cc countrycodes.Code) string {
	switch cc {
	case countrycodes.UK:
		r := randomizer.Roll1DN(len(country.SurnamesUK))
		return country.SurnamesUK[r-1]
	case countrycodes.Germany:
		r := randomizer.Roll1DN(len(country.SurnamesGermany))
		return country.SurnamesGermany[r-1]
	case countrycodes.USA:
		r := randomizer.Roll1DN(len(country.SurnamesUSA))
		return country.SurnamesUSA[r-1]
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
	case countrycodes.USA:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesUSA))
		return country.MaleFirstNamesUSA[r-1]
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
	case countrycodes.USA:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesUSA))
		return country.FemaleFirstNamesUSA[r-1]
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
		r := randomizer.Roll1DN(len(country.MaleFirstNamesUSA))
		firstname := strings.Title(strings.ToLower(country.MaleFirstNamesUSA[r-1]))
		r = randomizer.Roll1DN(len(country.SurnamesUSA))
		surname := strings.Title(strings.ToLower(country.SurnamesUSA[r-1]))
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
		surname := strings.Title(country.SurnamesUK[r-1])
		return firstname + " " + surname
	case countrycodes.Germany:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesGermany))
		firstname := country.FemaleFirstNamesGermany[r-1]
		r = randomizer.Roll1DN(len(country.SurnamesGermany))
		surname := strings.Title(country.SurnamesGermany[r-1])
		return firstname + " " + surname
	case countrycodes.USA:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesUSA))
		firstname := strings.Title(strings.ToLower(country.FemaleFirstNamesUSA[r-1]))
		r = randomizer.Roll1DN(len(country.SurnamesUSA))
		surname := strings.Title(strings.ToLower(country.SurnamesUSA[r-1]))
		return firstname + " " + surname
	}
	return countrycodes.InvalidParameter
}
