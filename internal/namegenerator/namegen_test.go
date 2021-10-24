package namegenerator

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator/country"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSurname(t *testing.T) {
	randomizer.Init()
	sn := CreateSurname(countrycodes.UK)
	fmt.Println(sn)
	assert.NotEqual(t, countrycodes.InvalidParameter, sn)
}

func TestFullname(t *testing.T) {
	randomizer.Init()
	fn := CreateMaleFullName(countrycodes.UK)
	fmt.Println(fn)
	fn = CreateFemaleFullName(countrycodes.UK)
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
	dup, found := findDuplicates(country.SurnamesUK)
	if found {
		t.Log("duplicate found: SurnamesUK: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(country.MaleFirstNamesUK)
	if found {
		t.Log("duplicate found: MaleFirstNamesUK: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(country.FemaleFirstNamesUK)
	if found {
		t.Log("duplicate found: FemaleFirstNamesUK: " + dup)
		t.FailNow()
	}

}