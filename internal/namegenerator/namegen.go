package namegenerator

import (
	"github.com/CmdSoda/boardgamewars/internal/namegenerator/languages"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

type LanguageCode int

const (
	EnglishAndUS     LanguageCode = 0
	German           LanguageCode = 1
	InvalidParameter string       = "InvalidParameter"
)

func CreateSurname(cc LanguageCode) string {
	switch cc {
	case EnglishAndUS:
		r := randomizer.Roll1DN(len(languages.EnglishSurnames))
		return languages.EnglishSurnames[r-1]
	case German:
		r := randomizer.Roll1DN(len(languages.GermanSurnames))
		return languages.GermanSurnames[r-1]
	}
	return InvalidParameter
}

func CreateMaleFirstname(cc LanguageCode) string {
	switch cc {
	case EnglishAndUS:
		r := randomizer.Roll1DN(len(languages.EnglishMaleFirstNames))
		return languages.EnglishMaleFirstNames[r-1]
	case German:
		r := randomizer.Roll1DN(len(languages.GermanMaleFirstNames))
		return languages.GermanMaleFirstNames[r-1]
	}
	return InvalidParameter
}

func CreateFemaleFirstname(cc LanguageCode) string {
	switch cc {
	case EnglishAndUS:
		r := randomizer.Roll1DN(len(languages.EnglishFemaleFirstNames))
		return languages.EnglishFemaleFirstNames[r-1]
	case German:
		r := randomizer.Roll1DN(len(languages.GermanFemaleFirstNames))
		return languages.GermanFemaleFirstNames[r-1]
	}
	return InvalidParameter
}

func CreateMaleFullName(cc LanguageCode) string {
	switch cc {
	case EnglishAndUS:
		r := randomizer.Roll1DN(len(languages.EnglishMaleFirstNames))
		firstname := languages.EnglishMaleFirstNames[r-1]
		r = randomizer.Roll1DN(len(languages.EnglishSurnames))
		surname := languages.EnglishSurnames[r-1]
		return firstname + " " + surname
	case German:
		r := randomizer.Roll1DN(len(languages.GermanMaleFirstNames))
		firstname := languages.GermanMaleFirstNames[r-1]
		r = randomizer.Roll1DN(len(languages.GermanSurnames))
		surname := languages.GermanSurnames[r-1]
		return firstname + " " + surname
	}
	return InvalidParameter
}
