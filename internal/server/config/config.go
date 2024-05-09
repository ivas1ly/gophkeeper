package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	defaultAddress              = "localhost:8080"
	defaultLogLevel             = "info"
	defaultDatabaseConnTimeout  = 5 * time.Second
	defaultDatabaseConnAttempts = 3
)

type Config struct {
	App
	Server
	DB
}

type App struct {
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

type Server struct {
	RunAddress string `mapstructure:"SERVER_ADDRESS"`
}

type DB struct {
	DatabaseURI          string        `mapstructure:"DB_ADDRESS"`
	DatabaseConnTimeout  time.Duration `mapstructure:"DB_CONN_TIMEOUT"`
	DatabaseConnAttempts int           `mapstructure:"DB_CONN_ATTEMPTS"`
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
	viper.SetDefault("db.address", "")
	viper.SetDefault("db.conn.timeout", defaultDatabaseConnTimeout)
	viper.SetDefault("db.conn.attempts", defaultDatabaseConnAttempts)

	cfg := Config{
		App: App{
			LogLevel: viper.GetString("log_level"),
		},
		Server: Server{
			RunAddress: viper.GetString("server.address"),
		},
		DB: DB{
			DatabaseURI:          viper.GetString("db.address"),
			DatabaseConnTimeout:  viper.GetDuration("db.conn.timeout"),
			DatabaseConnAttempts: viper.GetInt("db.conn.attempts"),
		},
	}

	fmt.Printf("\nstart application with config:\n%+v\n\n", cfg)

	return cfg, nil
}
