package main

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config AppConfig

type AppConfig struct {
	ServiceName      string `mapstructure:"SERVICE_NAME"`
	Port             string `mapstructure:"PORT"`
	IsJsonLogging    bool   `mapstructure:"JSON_LOGGING"`
	IsLogLevelDebug  bool   `mapstructure:"LOG_LEVEL_DEBUG"`
	IsTracingEnabled bool   `mapstructure:"TRACING_ENABLED"`
	IsDevMode        bool   `mapstructure:"DEV_MODE"`
	JaegerEndpoint   string `mapstructure:"JAEGER_ENDPOINT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostresPassword  string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	JwkSetUri        string `mapstructure:"JWK_SET_URI"`
}

func SetupConfig() {
	viper.SetDefault("SERVICE_NAME", "cng-hello-backend")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("JSON_LOGGING", true)
	viper.SetDefault("LOG_LEVEL_DEBUG", false)
	viper.SetDefault("TRACING_ENABLED", true)
	viper.SetDefault("DEV_MODE", false)
	viper.SetDefault("JAEGER_ENDPOINT", "")
	viper.SetDefault("POSTGRES_HOST", "")
	viper.SetDefault("POSTGRES_PORT", "5432")
	viper.SetDefault("POSTGRES_USER", "")
	viper.SetDefault("POSTGRES_PASSWORD", "")
	viper.SetDefault("POSTGRES_DB", "")
	viper.SetDefault("JWK_SET_URI", "")

	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.ReadInConfig()

	err := viper.Unmarshal(&Config)
	if err != nil {
		fmt.Println("Could not unmarshal config")
	}

	if Config.IsDevMode {
		fmt.Printf("%+v\n", Config)
	}

}
