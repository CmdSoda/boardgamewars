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
	out :=s[i1+1:i2]
	return out
}

type NameSet struct {
	Males []string
	Females []string
	Surnames []string
}

var theNameSet = map[countrycodes.Code]NameSet{
	countrycodes.UK: {
		Males:    country.MaleFirstNamesUK,
		Females:  country.FemaleFirstNamesUK,
		Surnames: country.SurnamesUK,
	},
	countrycodes.Germany: {
		Males:    country.MaleFirstNamesGermany,
		Females:  country.FemaleFirstNamesGermany,
		Surnames: country.SurnamesGermany,
	},
	countrycodes.USA: {
		Males:    country.MaleFirstNamesUSA,
		Females:  country.FemaleFirstNamesUK,
		Surnames: country.SurnamesUK,
	},
	countrycodes.Russia: {
		Males:    country.MaleFirstNamesRussia,
		Females:  country.FemaleFirstNamesRussia,
		Surnames: country.SurnamesRussia,
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
