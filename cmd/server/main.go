package main

import (
	"context"
	"log"

	server "github.com/ivas1ly/gophkeeper/internal/server/app"
	"github.com/ivas1ly/gophkeeper/internal/server/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		log.Printf("unable to load config: %s", err.Error())
		return
	}

	s, err := server.New(ctx, cfg)
	if err != nil {
		log.Printf("can't init server: %s", err.Error())
		return
	}

	if err = s.Run(ctx); err != nil {
		log.Printf("server terminated with error: %s", err.Error())
	}
}
