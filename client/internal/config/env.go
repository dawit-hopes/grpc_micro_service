// Package config env configures the env variables
package config

import (
	"github.com/caarlos0/env/v10"
	"log"
)

type Env struct {
	ServerAddress string `env:"SERVER_ADDRESS,required"`
	Serverhost    string `env:"SERVER_HOST,required"`
}

func NewEnv() *Env {
	cfg := &Env{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("❌ Failed to load environment variables: %v", err)
	}
	return cfg
}
