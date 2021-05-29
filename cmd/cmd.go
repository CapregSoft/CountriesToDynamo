package cmd

import (
	"github.com/CapregSoft/CountriesToDynamo/pkg/config"
	"github.com/CapregSoft/CountriesToDynamo/pkg/dynamo"
	util "github.com/CapregSoft/CountriesToDynamo/pkg/utils"
)

// Starting point of the cmd
func Start(file string) {
	// TODO : take the file from the command line
	err := config.Load()
	if err != nil {
		panic(err)
	}
	countries := util.ReadCountries(file)
	countriesRecord := util.CountriesToModel(countries)
	dynamo.Insert(countriesRecord)
}
