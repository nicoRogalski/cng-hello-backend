package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var App Config

type Config struct {
	ServiceName     string `mapstructure:"SERVICE_NAME"`
	Port            string `mapstructure:"PORT"`
	IsJsonLogging   bool   `mapstructure:"JSON_LOGGING"`
	IsLogLevelDebug bool   `mapstructure:"LOG_LEVEL_DEBUG"`
	IsDevMode       bool   `mapstructure:"DEV_MODE"`
	JwtCertUrl      string `mapstructure:"JWT_CERT_URL"`
	JwtSecret       string `mapstructure:"JWT_SECRET"`
	JaegerEndpoint  string `mapstructure:"JAEGER_ENDPOINT"`
	PostgresHost    string `mapstructure:"POSTGRES_HOST"`
	PostgresUser    string `mapstructure:"POSTGRES_USER"`
	PostresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb      string `mapstructure:"POSTGRES_DB"`
}

func init() {
	viper.SetDefault("SERVICE_NAME", "cng-hello-backend")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("JSON_LOGGING", true)
	viper.SetDefault("LOG_LEVEL_DEBUG", false)
	viper.SetDefault("DEV_MODE", false)
	viper.SetDefault("JAEGER_ENDPOINT", "")
	viper.SetDefault("POSTGRES_HOST", "")
	viper.SetDefault("POSTGRES_USER", "")
	viper.SetDefault("POSTGRES_PASSWORD", "")
	viper.SetDefault("POSTGRES_DB", "")

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

	viper.AutomaticEnv()

	err := viper.Unmarshal(&App)
	if err != nil {
		panic("Could not unmarshal config")
	}
}
