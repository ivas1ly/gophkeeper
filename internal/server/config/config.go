package config

import (
	"fmt"
	"net"
)

const (
	defaultHost     = "localhost"
	defaultPort     = "8080"
	defaultLogLevel = "info"
)

type Config struct {
	Server
	HTTP
}

type Server struct {
	LogLevel string
}

type HTTP struct {
	RunAddress string
}

func New() Config {
	cfg := Config{
		Server: Server{
			LogLevel: defaultLogLevel,
		},
		HTTP: HTTP{
			RunAddress: net.JoinHostPort(defaultHost, defaultPort),
		},
	}

	fmt.Printf("\nstart application with config:\n%+v\n\n", cfg)

	return cfg
}
