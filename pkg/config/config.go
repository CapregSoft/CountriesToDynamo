package config

import (
	"github.com/spf13/viper"
)

// Config is a global variable that holds all the applications config
var configs Config

// Config holds all applications configs
type Config struct {
	Region string `mapstructure:"region"`
}

// Load configs from ./config/ yml files depending on APP_ENV.
func Load() error {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetEnvPrefix("")
	v.AutomaticEnv()

	v.SetConfigType("yaml")

	v.SetConfigName("env/default")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.Unmarshal(&configs)
	return nil
}

// Get returns previously loaded configs
func GetConfig() Config {
	return configs
}
