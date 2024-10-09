package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config

	viper.SetConfigName("main_config")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error reading config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("fatal error parsing the .env file: %w", err)
	}

	return &config, nil
}
