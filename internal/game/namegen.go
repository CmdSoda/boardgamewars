package game

import (
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
	Country       string
	Males         []string
	Females       []string
	Surnames      []string
	Cities        []string
	AirForceBases []string
}

type Generator map[string]*NameSet

func (g *Generator) AddNameSet(ns *NameSet) {
	tns := *g
	tns[ns.Country] = ns
}

func (g *Generator) CreateMaleFullName(countryName string) string {
	tns := *g
	r := randomizer.Roll1DN(len(tns[countryName].Males))
	firstname := ignoreBraces(strings.Title(strings.ToLower(tns[countryName].Males[r-1])))
	r = randomizer.Roll1DN(len(tns[countryName].Surnames))
	surname := ignoreBraces(strings.Title(strings.ToLower(tns[countryName].Surnames[r-1])))
	return firstname + " " + surname
}

func (g *Generator) CreateFemaleFullName(countryName string) string {
	tns := *g
	r := randomizer.Roll1DN(len(tns[countryName].Females))
	firstname := ignoreBraces(strings.Title(strings.ToLower(tns[countryName].Females[r-1])))
	r = randomizer.Roll1DN(len(tns[countryName].Surnames))
	surname := ignoreBraces(strings.Title(strings.ToLower(tns[countryName].Surnames[r-1])))
	return firstname + " " + surname
}

func (g *Generator) CreateFullName(male bool, countryName string) string {
	if male {
		return g.CreateMaleFullName(countryName)
	} else {
		return g.CreateFemaleFullName(countryName)
	}
}

func (g *Generator) CreateCityName(countryName string) string {
	tns := *g
	r := randomizer.Roll1DN(len(tns[countryName].Cities))
	city := ignoreBraces(strings.Title(strings.ToLower(tns[countryName].Cities[r-1])))
	return city
}

func (g *Generator) CreateAirForceBaseName(countryName string) string {
	tns := *g
	r := randomizer.Roll1DN(len(tns[countryName].AirForceBases))
	airforcebase := ignoreBraces(strings.Title(strings.ToLower(tns[countryName].AirForceBases[r-1])))
	return airforcebase
}
