package config

import (
	"github.com/joho/godotenv"
	"fmt"
	"log"
	"os"
)

type fromEnvConfig struct {
	DatabaseUrl string
	HostPort string
}

type EnvironmentError struct {
	Message string
}

func (e *EnvironmentError) Error() string {
	return e.Message
}

func loadEnvFile(envFilePath string) {
	if envFilePath != "" {
		err := godotenv.Load(envFilePath)
		if err != nil {
			log.Printf("%s file found. This might've been a mistake. Continuing without it.", envFilePath)
		}
	}
}

func loadFromEnv(envFilePath string) (fromEnvConfig, error) {
	loadEnvFile(envFilePath)

	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	postgresEnv := map[string]string{
		"POSTGRES_USER": os.Getenv("POSTGRES_USER"),
		"POSTGRES_PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"POSTGRES_HOST": os.Getenv("POSTGRES_HOST"),
		"POSTGRES_DB": os.Getenv("POSTGRES_DB"),
	};

	for env, value := range postgresEnv {
		if value == "" {
			return fromEnvConfig{}, &EnvironmentError{Message: fmt.Sprintf("Environment variable %s is not set", env)}
		}
	}
	
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		return fromEnvConfig{}, &EnvironmentError{
			Message: "Environment variable BACKEND_PORT is not set. You must set at least the port for the backend to run",
		}
	}

	return fromEnvConfig{ 
		DatabaseUrl: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			postgresEnv["POSTGRES_USER"],
			postgresEnv["POSTGRES_PASSWORD"],
			postgresEnv["POSTGRES_HOST"],
			dbPort,
			postgresEnv["POSTGRES_DB"],
		),
		HostPort: fmt.Sprintf(
			"%s:%s",
			os.Getenv("BACKEND_HOST"),
			port,
		),
	}, nil
}
