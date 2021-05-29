package util

import (
	"encoding/json"
	"io/ioutil"
)

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
