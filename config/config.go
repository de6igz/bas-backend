package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"db"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Path string `mapstructure:"path"`
}

func LoadConfig() (*Config, error) {
	var config Config

	env := os.Getenv("CONFIG_ENV")
	if env == "" {
		env = "yaml"
	}

	if env == "env" {
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// Explicitly set default values to ensure that Viper can bind them
		viper.SetDefault("db.path", "./db/data.sqlite")
		viper.SetDefault("server.port", 8080)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
