// Package config implements the struct to read envs
package config

import (
	"github.com/caarlos0/env/v10"
	"log"
)

type Env struct {
	MongoDBConnectionString string `env:"MONGODB_CONNECTION_STRING,required"`
}

func NewEnv() *Env {
	cfg := &Env{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("❌ Failed to load environment variables: %v", err)
	}
	return cfg
}
