package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName      string `mapstructure:"SERVICE_NAME"`
	Port             string `mapstructure:"PORT"`
	IsJsonLogging    bool   `mapstructure:"JSON_LOGGING"`
	LogLevel         string `mapstructure:"LOG_LEVEL"`
	IsTracingEnabled bool   `mapstructure:"TRACING_ENABLED"`
	IsMetricsEnabled bool   `mapstructure:"METRICS_ENABLED"`
	IsDevMode        bool   `mapstructure:"DEV_MODE"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostresPassword  string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	JwkSetUri        string `mapstructure:"JWK_SET_URI"`
}

func Load() *Config {
	viper.SetDefault("SERVICE_NAME", "cng-hello-backend")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("JSON_LOGGING", true)
	viper.SetDefault("LOG_LEVEL", "INFO")
	viper.SetDefault("TRACING_ENABLED", true)
	viper.SetDefault("METRICS_ENABLED", true)
	viper.SetDefault("DEV_MODE", false)
	viper.SetDefault("POSTGRES_HOST", "")
	viper.SetDefault("POSTGRES_PORT", "5432")
	viper.SetDefault("POSTGRES_USER", "")
	viper.SetDefault("POSTGRES_PASSWORD", "")
	viper.SetDefault("POSTGRES_DB", "")
	viper.SetDefault("JWK_SET_URI", "")

	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	cfg := new(Config)
	err := viper.Unmarshal(cfg)
	if err != nil {
		panic("Could not unmarshal config")
	}
	return cfg
}
