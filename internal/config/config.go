package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Viper *viper.Viper
}

var Viper *viper.Viper

// Init initializes viper configuration and returns a Config instance.
func Init() (*Config, error) {
	Viper = viper.New()
	Viper.SetConfigName("meld")
	Viper.SetConfigType("yaml")
	Viper.AddConfigPath(".")

	// Default configuration values.
	Viper.SetDefault("version", "0.1.0-dev")
	Viper.SetDefault("verbose", false)

	cfg := &Config{Viper: Viper}
	return cfg, nil
}
