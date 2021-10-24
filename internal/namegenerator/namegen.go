package namegenerator

import (
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator/country"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"strings"
)

func IgnoreBraces(s string) string {
	i := strings.Index(s, "(")
	if i == -1 {
		return s
	}
	return s[:i-1]
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
	case countrycodes.Russia:
		r := randomizer.Roll1DN(len(country.MaleFirstNamesRussia))
		firstname := IgnoreBraces(country.MaleFirstNamesRussia[r-1])
		r = randomizer.Roll1DN(len(country.SurnamesRussia))
		surname := IgnoreBraces(country.SurnamesRussia[r-1])
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
	case countrycodes.Russia:
		r := randomizer.Roll1DN(len(country.FemaleFirstNamesRussia))
		firstname := IgnoreBraces(country.FemaleFirstNamesRussia[r-1])
		r = randomizer.Roll1DN(len(country.SurnamesRussia))
		surname := IgnoreBraces(country.SurnamesRussia[r-1])
		return firstname + " " + surname
	}
	return countrycodes.InvalidParameter
}
