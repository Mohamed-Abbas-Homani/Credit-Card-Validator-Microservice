package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port           int             `mapstructure:"PORT"`
	GRPCPort       int             `mapstructure:"GRPC_PORT"`
	LogLevel       string          `mapstructure:"LOG_LEVEL"`
	MetricsEnabled bool            `mapstructure:"METRICS_ENABLED"`
	Validator      ValidatorConfig `mapstructure:",squash"`
}

type ValidatorConfig struct {
	EnableBINLookup bool          `mapstructure:"ENABLE_BIN_LOOKUP"`
	HTTPTimeout     time.Duration `mapstructure:"HTTP_TIMEOUT"`
	BINServiceURL   string        `mapstructure:"BIN_SERVICE_URL"`
	MaskSensitive   bool          `mapstructure:"MASK_SENSITIVE"`
}

// Load returns merged service and validator configuration
func Load() *Config {
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("GRPC_PORT", 9090)
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("METRICS_ENABLED", true)

	viper.SetDefault("ENABLE_BIN_LOOKUP", true)
	viper.SetDefault("HTTP_TIMEOUT", "10s")
	viper.SetDefault("BIN_SERVICE_URL", "https://lookup.binlist.net")
	viper.SetDefault("MASK_SENSITIVE", true)

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
