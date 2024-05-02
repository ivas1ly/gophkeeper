package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultAddress  = "localhost:8080"
	defaultLogLevel = "info"
)

type Config struct {
	App
	Server
}

type App struct {
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

type Server struct {
	RunAddress string `mapstructure:"SERVER_ADDRESS"`
}

func Load() (Config, error) {
	viper.SetEnvPrefix("GK")
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("fatal error read config file: %w", err)
	}

	viper.SetDefault("server.address", defaultAddress)
	viper.SetDefault("log_level", defaultLogLevel)

	cfg := Config{
		App: App{
			LogLevel: viper.GetString("log_level"),
		},
		Server: Server{
			RunAddress: viper.GetString("server.address"),
		},
	}

	fmt.Printf("\nstart application with config:\n%+v\n\n", cfg)

	return cfg, nil
}
