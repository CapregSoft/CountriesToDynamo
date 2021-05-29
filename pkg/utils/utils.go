package util

import (
	"encoding/json"
	"io/ioutil"

	"github.com/CapregSoft/CountriesToDynamo/pkg/model"
)

// Read the countries from the file and convert it into map
func ReadCountries(file string) map[string][]string {
	var countries map[string][]string
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	// Unmarshal data
	err = json.Unmarshal(data, &countries)
	if err != nil {
		panic(err)
	}
	return countries
}

// CountriesToModel converts the map into json values
func CountriesToModel(countries map[string][]string) []model.Country {
	countryJson := make([]model.Country, len(countries))
	i := 0
	for k, v := range countries {
		countryJson[i] = model.Country{
			CountryName: k,
			Cities:      v,
		}
		i++
	}
	return countryJson
}
