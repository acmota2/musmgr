package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	platform "github.com/acmota2/musmgr/backend/internal/platform/file_access"
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
	StorageConfig  *platform.StorageConfig
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

func requireEnvs(scope string, keys ...string) (map[string]string, error) {
	envMap := make(map[string]string, len(keys))
	for _, k := range keys {
		env := os.Getenv(k)
		if env == "" {
			return nil, fmt.Errorf("%s: %s must be defined", scope, k)
		}
	}

	return envMap, nil
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

func validateStorageConfig(storageType platform.StorageType) (*platform.StorageConfig, error) {
	switch storageType {
	case platform.LOCAL:
		if path := os.Getenv("LOCAL_PATH"); path != "" {
			return &platform.StorageConfig{
				LocalPath: path,
			}, nil
		} else {
			return nil, fmt.Errorf("LOCAL_PATH must be defined for storage type LOCAL")
		}
	case platform.MINIO:
		minioSettings, err := requireEnvs(
			"MinIO",
			"MINIO_ENDPOINT",
			"MINIO_ACCESS_KEY",
			"MINIO_SECRET_KEY",
			"MINIO_BUCKET_NAME",
			"MINIO_BUCKET_REGION",
			"MINIO_SSL",
		)
		if err != nil {
			return nil, err
		}

		return &platform.StorageConfig{
			MinioEndpoint:        minioSettings["MINIO_ENDPOINT"],
			MinioAccessKeyId:     minioSettings["MINIO_ACCESS_KEY"],
			MinioSecretAccessKey: minioSettings["MINIO_SECRET_KEY"],
			MinioBucketName:      minioSettings["MINIO_BUCKET_NAME"],
			MinioBucketRegion:    minioSettings["MINIO_BUCKET_REGION"],
			MinioSSL:             minioSettings["MINIO_SSL"] == "true",
		}, nil
	default:
		return nil, fmt.Errorf("Invalid storage type")
	}
}

func loadFromEnv(argsConfig configFromArgs) (*fromEnvConfig, error) {
	loadEnvFile(argsConfig.EnvFilePath)

	postgresEnv, err := requireEnvs(
		"Database",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_HOST",
		"POSTGRES_DB",
	)
	if err != nil {
		return nil, err
	}

	storageConfig, err := validateStorageConfig(argsConfig.StorageType)
	if err != nil {
		return nil, err
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
		StorageConfig:  storageConfig,
	}, nil
}
