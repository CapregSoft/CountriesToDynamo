package cmd

import (
	"github.com/CapregSoft/CountriesToDynamo/pkg/config"
	util "github.com/CapregSoft/CountriesToDynamo/pkg/utils"
)

func Start(file string) {
	err := config.Load()
	if err != nil {
		panic(err)
	}
	util.ReadCountries(file)
}
