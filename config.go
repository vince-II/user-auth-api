package main

import (
	"os"
)

type Configuration struct {
	Port        string
	DatabaseURL string
}

func NewConfiguration() *Configuration {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
	}

	return &Configuration{
		DatabaseURL: "postgres://user:password@localhost:5432/mydb",
	}
}
