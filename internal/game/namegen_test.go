package game

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	assert.Nil(t, InitGame(0))

	for _, item := range Globals.CountryDataMap {
		dup, found := findDuplicates(item.NameSet.Surnames)
		if found {
			t.Log("duplicate found: surnames in " + item.Country + ": " + dup)
			t.FailNow()
		}
		dup, found = findDuplicates(item.NameSet.Males)
		if found {
			t.Log("duplicate found: males in " + item.Country + ": " + dup)
			t.FailNow()
		}
		dup, found = findDuplicates(item.NameSet.Females)
		if found {
			t.Log("duplicate found: females in " + item.Country + ": " + dup)
			t.FailNow()
		}
		dup, found = findDuplicates(item.NameSet.Cities)
		if found {
			t.Log("duplicate found: cities in " + item.Country + ": " + dup)
			t.FailNow()
		}
		dup, found = findDuplicates(item.NameSet.AirForceBases)
		if found {
			t.Log("duplicate found: air bases in " + item.Country + ": " + dup)
			t.FailNow()
		}
	}
}

func TestIgnoreBraces(t *testing.T) {
	is := ignoreBraces("bla (uschi)")
	fmt.Println(is)
}
