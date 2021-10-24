package namegenerator

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator/languages"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSurname(t *testing.T) {
	randomizer.Init()
	sn := CreateSurname(EnglishAndUS)
	fmt.Println(sn)
	assert.NotEqual(t, InvalidParameter, sn)
}

func TestFullname(t *testing.T) {
	randomizer.Init()
	fn := CreateMaleFullName(EnglishAndUS)
	fmt.Println(fn)
}

func findDuplicates(a []string) (string, bool) {
	for i := 0; i < len(a); i++ {
		first := a[i]
		for j := 0; j < len(a); j++ {
			compare := a[j]
			if i != j {
				if first == compare {
					return compare, true
				}
			}
		}
	}
	return "", false
}

func TestFindDuplicates(t *testing.T) {
	dup, found := findDuplicates(languages.EnglishSurnames)
	if found {
		t.Log("duplicate found: EnglishSurnames: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(languages.EnglishMaleFirstNames)
	if found {
		t.Log("duplicate found: EnglishMaleFirstNames: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(languages.EnglishFemaleFirstNames)
	if found {
		t.Log("duplicate found: EnglishFemaleFirstNames: " + dup)
		t.FailNow()
	}

}