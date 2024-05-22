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
	defaultAppName              = "GophKeeper"
	defaultDatabaseConnTimeout  = 5 * time.Second
	defaultDatabaseConnAttempts = 3
	defaultExpirationTime       = 12 * time.Hour
)

var (
	defaultSigningKey = []byte("30ab73e47d6aa53116b90aa23bc3b9ae68822d10f7288ba10f2d901879539bde")
)

type Config struct {
	App
	Server
	DB
}

type App struct {
	LogLevel       string        `mapstructure:"APP_LOG_LEVEL"`
	Name           string        `mapstructure:"APP_NAME"`
	SigningKey     []byte        `mapstructure:"APP_JWT_SIGNING_KEY"`
	ExpirationTime time.Duration `mapstructure:"APP_JWT_EXPIRATION_TIME"`
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
	viper.SetDefault("app.name", defaultAppName)
	viper.SetDefault("app.jwt.signing_key", defaultSigningKey)
	viper.SetDefault("app.jwt.expiration_time", defaultExpirationTime)

	cfg := Config{
		App: App{
			LogLevel:       viper.GetString("app.log_level"),
			Name:           viper.GetString("app.name"),
			SigningKey:     []byte(viper.GetString("app.jwt.signing_key")),
			ExpirationTime: viper.GetDuration("app.jwt.expiration_time"),
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
