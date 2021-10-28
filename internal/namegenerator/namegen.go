package namegenerator

import (
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator/country"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"strings"
)

func ignoreBraces(s string) string {
	i1 := strings.Index(s, "(")
	i2 := strings.Index(s, ")")
	if i1 == -1 {
		return s
	}
	out := s[i1+1 : i2]
	return out
}

type NameSet struct {
	Males        []string
	Females      []string
	Surnames     []string
	Cities       []string
	AirForceBase []string
}

var theNameSet = map[countrycodes.Code]NameSet{
	countrycodes.UK: {
		Males:        country.MaleFirstNamesUK,
		Females:      country.FemaleFirstNamesUK,
		Surnames:     country.SurnamesUK,
		Cities:       country.CitiesUK,
		AirForceBase: country.AirForceBasesUK,
	},
	countrycodes.Germany: {
		Males:        country.MaleFirstNamesGermany,
		Females:      country.FemaleFirstNamesGermany,
		Surnames:     country.SurnamesGermany,
		Cities:       country.CitiesGermany,
		AirForceBase: country.AirForceBasesGermany,
	},
	countrycodes.USA: {
		Males:        country.MaleFirstNamesUSA,
		Females:      country.FemaleFirstNamesUSA,
		Surnames:     country.SurnamesUSA,
		Cities:       country.CitiesUSA,
		AirForceBase: country.AirForceBasesUSA,
	},
	countrycodes.Russia: {
		Males:        country.MaleFirstNamesRussia,
		Females:      country.FemaleFirstNamesRussia,
		Surnames:     country.SurnamesRussia,
		Cities:       country.CitiesRussia,
		AirForceBase: country.AirForceBasesRussia,
	},
}

func CreateMaleFullName(cc countrycodes.Code) string {
	r := randomizer.Roll1DN(len(theNameSet[cc].Males))
	firstname := ignoreBraces(strings.Title(strings.ToLower(theNameSet[cc].Males[r-1])))
	r = randomizer.Roll1DN(len(theNameSet[cc].Surnames))
	surname := ignoreBraces(strings.Title(strings.ToLower(theNameSet[cc].Surnames[r-1])))
	return firstname + " " + surname
}

func CreateFemaleFullName(cc countrycodes.Code) string {
	r := randomizer.Roll1DN(len(theNameSet[cc].Females))
	firstname := ignoreBraces(strings.Title(strings.ToLower(theNameSet[cc].Females[r-1])))
	r = randomizer.Roll1DN(len(theNameSet[cc].Surnames))
	surname := ignoreBraces(strings.Title(strings.ToLower(theNameSet[cc].Surnames[r-1])))
	return firstname + " " + surname
}

func CreateCityName(cc countrycodes.Code) string {
	r := randomizer.Roll1DN(len(theNameSet[cc].Cities))
	city := ignoreBraces(strings.Title(strings.ToLower(theNameSet[cc].Cities[r-1])))
	return city
}

func CreateAirForceBaseName(cc countrycodes.Code) string {
	r := randomizer.Roll1DN(len(theNameSet[cc].AirForceBase))
	airforcebase := ignoreBraces(strings.Title(strings.ToLower(theNameSet[cc].AirForceBase[r-1])))
	return airforcebase
}
