package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	// Define your configuration variables here
	Port             string
	OpenaiServiceUrl string
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config") // Base configuration file name
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	// Load base config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Load environment specific config
	env := os.Getenv("APP_ENV")
	if env != "" {
		viper.SetConfigName("config." + strings.ToLower(env)) // e.g., config.production
		viper.MergeInConfig()                                 // This merges the specific config into the base config
	}

	viper.AutomaticEnv() // Override config values with environment variables

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
