package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	ServiceName     string `mapstructure:"SERVICE_NAME"`
	Port            string `mapstructure:"PORT"`
	IsJsonLogging   bool   `mapstructure:"JSON_LOGGING"`
	IsLogLevelDebug bool   `mapstructure:"LOG_LEVEL_DEBUG"`
	IsDevMode       bool   `mapstructure:"DEV_MODE"`
}

func init() {

	viper.SetDefault("SERVICE_NAME", "cng-hello-backend")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("JSON_LOGGING", true)
	viper.SetDefault("LOG_LEVEL_DEBUG", false)
	viper.SetDefault("DEV_MODE", false)

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
