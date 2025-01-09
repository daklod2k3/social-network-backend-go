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

func GetConfig() *Configuration {
	fmt.Println("init config")
	viper.SetConfigName("config")        // name of config file (without extension)
	viper.AddConfigPath("shared/config") // path to look for the config file in
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("parse config file error: %w", err))
	}
	return &config
}
