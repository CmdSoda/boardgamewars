package namegenerator

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator/country"
	"testing"
)

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

	dup, found = findDuplicates(country.SurnamesGermany)
	if found {
		t.Log("duplicate found: SurnamesUK: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(country.MaleFirstNamesGermany)
	if found {
		t.Log("duplicate found: MaleFirstNamesGermany: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(country.FemaleFirstNamesGermany)
	if found {
		t.Log("duplicate found: FemaleFirstNamesGermany: " + dup)
		t.FailNow()
	}

	//
	// USA
	//

	dup, found = findDuplicates(country.SurnamesUSA)
	if found {
		t.Log("duplicate found: SurnamesUSA: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(country.MaleFirstNamesUSA)
	if found {
		t.Log("duplicate found: MaleFirstNamesUSA: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(country.FemaleFirstNamesUSA)
	if found {
		t.Log("duplicate found: FemaleFirstNamesUSA: " + dup)
		t.FailNow()
	}

	//
	// Russia
	//

	dup, found = findDuplicates(country.SurnamesRussia)
	if found {
		t.Log("duplicate found: SurnamesRussia: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(country.MaleFirstNamesRussia)
	if found {
		t.Log("duplicate found: MaleFirstNamesRussia: " + dup)
		t.FailNow()
	}

	dup, found = findDuplicates(country.FemaleFirstNamesRussia)
	if found {
		t.Log("duplicate found: FemaleFirstNamesRussia: " + dup)
		t.FailNow()
	}

}

func TestIgnoreBraces(t *testing.T) {
	is := ignoreBraces("bla (uschi)")
	fmt.Println(is)
}
