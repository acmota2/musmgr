package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	Development DeploymentType = "DEVELOPMENT"
	Production  DeploymentType = "PRODUCTION"
)

type DeploymentType string

type fromEnvConfig struct {
	AdminRoutes    []string
	DatabaseUrl    string
	DeploymentMode string
	PublicRoutes   []string
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

func parseDeploymentMode(mode string) string {
	switch mode {
	case "DEVELOPMENT", "PRODUCTION":
		return mode
	default:
		log.Println("Invalid or empty DEPLOYMENT_TYPE. Defaulting to DEVELOPMENT")
	}
	return "DEVELOPMENT"
}

func parseAllowedOrigins(list string) ([]string, error) {
	var origins []string
	if err := json.Unmarshal([]byte(list), &origins); err != nil {
		return nil, err
	}
	return origins, nil
}

func loadFromEnv(envFilePath string) (*fromEnvConfig, error) {
	loadEnvFile(envFilePath)

	postgresEnv := map[string]string{
		"POSTGRES_USER":     os.Getenv("POSTGRES_USER"),
		"POSTGRES_PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"POSTGRES_HOST":     os.Getenv("POSTGRES_HOST"),
		"POSTGRES_DB":       os.Getenv("POSTGRES_DB"),
	}

	for env, value := range postgresEnv {
		if value == "" {
			return &fromEnvConfig{}, &EnvironmentError{Message: fmt.Sprintf("Environment variable %s is not set", env)}
		}
	}

	mode := parseDeploymentMode(os.Getenv("DEPLOYMENT_TYPE"))
	allowedPublicRoutes, err := parseAllowedOrigins(os.Getenv("ALLOWED_PUBLIC_DOMAINS"))
	if err != nil {
		return nil, err
	}
	allowedAdminRoutes, err := parseAllowedOrigins(os.Getenv("ALLOWED_ADMIN_DOMAINS"))
	if err != nil {
		return nil, err
	}

	return &fromEnvConfig{
		AdminRoutes: allowedAdminRoutes,
		DatabaseUrl: fmt.Sprintf(
			"postgres://%s:%s@%s/%s",
			postgresEnv["POSTGRES_USER"],
			postgresEnv["POSTGRES_PASSWORD"],
			postgresEnv["POSTGRES_HOST"],
			postgresEnv["POSTGRES_DB"],
		),
		DeploymentMode: mode,
		PublicRoutes:   allowedPublicRoutes,
	}, nil
}
