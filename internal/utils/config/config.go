package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	DevMode     bool   `mapstructure:"DEV_MODE"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
	Port        string `mapstructure:"PORT"`
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			log.Warn().Msg("Failed to load from config file")
		}
	}

	//TODO: Does not work with env variable ..os.Getenv is working..
	viper.AutomaticEnv()

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		panic("Could not unmarshal config")
	}
}
