package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct{}

func ReadConfig(path string) {
	viper.SetConfigFile("campus_forum")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func (c *Config) Domain() interface{} {
	return viper.Get("domain")
}
