package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port           int    `mapstructure:"PORT"`
	GRPCPort       int    `mapstructure:"GRPC_PORT"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	MetricsEnabled bool   `mapstructure:"METRICS_ENABLED"`
}

func Load() *Config {
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("GRPC_PORT", 9090)
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("METRICS_ENABLED", true)

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
