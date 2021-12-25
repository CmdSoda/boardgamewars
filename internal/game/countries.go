package game

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type CountryFilenames []string

type CountryName string

type CountryDataItem struct {
	Country        string
	FlightRankList []FlightRank
	NameSet        NameSet
}
type CountryDataMap map[CountryName]*CountryDataItem

func LoadCountryData() error {
	var err error
	countriesFilename := Globals.Startup.DataPath + "countries.json"
	var files []string
	file, err := os.Open(countriesFilename)
	if err != nil {
		Log.Errorf("%s not found\n", countriesFilename)
		return err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		Log.Errorf("%s while reading\n", countriesFilename)
		return err
	}
	err = json.Unmarshal(bytes, &files)
	if err != nil {
		Log.Errorf("%s while unmarshaling\n", countriesFilename)
		return err
	}

	Globals.CountryDataMap = map[CountryName]*CountryDataItem{}
	for _, filename := range files {
		cd := CountryDataItem{}
		countryDataFilename := Globals.Startup.DataPath + filename
		file, err := os.Open(countryDataFilename)
		if err != nil {
			Log.Errorf("%s not found\n", countryDataFilename)
			return err
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			Log.Errorf("%s while reading\n", countryDataFilename)
			return err
		}
		err = json.Unmarshal(bytes, &cd)
		if err != nil {
			Log.Errorf("%s while unmarshaling\n", countryDataFilename)
			return err
		}
		cd.NameSet.Country = cd.Country
		Globals.CountryDataMap[CountryName(cd.Country)] = &cd
	}

	return nil
}
