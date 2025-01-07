package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	Port     int
	Database struct {
		ConnectString string `mapstructure:"ConnectString"`
	}
	Supabase struct {
		Url string
		Key string
	}
}

var Config viper.Viper

func init() {
	Config := viper.New()
	Config.SetConfigName("config")        // name of config file (without extension)
	Config.AddConfigPath("shared/config") // path to look for the config file in
	// Find and read the config file
	if err := Config.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
