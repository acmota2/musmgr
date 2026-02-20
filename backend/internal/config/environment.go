package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	platform "github.com/acmota2/musmgr/backend/internal/platform/file-access"
)

const (
	Development DeploymentType = "DEVELOPMENT"
	Production  DeploymentType = "PRODUCTION"
)

type DeploymentType string

type EnvConfig struct {
	AdminRoutes          []string
	DatabaseUrl          string
	DeploymentMode       string
	PublicRoutes         []string
	LocalPath            string
	MinioEndpoint        string
	MinioBucketName      string
	MinioBucketRegion    string
	MinioAccessKeyId     string
	MinioSecretAccessKey string
	MinioSSL             bool
}

type EnvironmentError struct {
	Message string
}

func (e *EnvironmentError) Error() string {
	return e.Message
}

func requireEnvs(scope string, keys ...string) (map[string]string, error) {
	envMap := make(map[string]string, len(keys))
	for _, k := range keys {
		env := os.Getenv(k)
		if env == "" {
			return nil, fmt.Errorf("%s: %s must be defined", scope, k)
		}
		envMap[k] = env
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

func validateStorageConfig(storageType platform.StorageType) (map[string]string, error) {
	switch storageType {
	case platform.LOCAL:
		if path := os.Getenv("LOCAL_PATH"); path != "" {
			return map[string]string{
				"LOCAL_PATH": path,
			}, nil
		} else {
			return nil, fmt.Errorf("LOCAL_PATH must be defined for storage type LOCAL")
		}
	case platform.MINIO:
		minioSettings, err := requireEnvs(
			"MinIO",
			"MINIO_ENDPOINT",
			"MINIO_ACCESS_KEY_ID",
			"MINIO_SECRET_ACCESS_KEY",
			"MINIO_BUCKET_NAME",
			"MINIO_BUCKET_REGION",
			"MINIO_SSL",
		)
		if err != nil {
			return nil, err
		}

		return minioSettings, nil
	default:
		return nil, fmt.Errorf("Invalid storage type")
	}
}

func LoadFromEnv(storageType platform.StorageType) (*EnvConfig, error) {
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

	storageEnvs, err := validateStorageConfig(storageType)
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

	return &EnvConfig{
		AdminRoutes: allowedAdminRoutes,
		DatabaseUrl: fmt.Sprintf(
			"postgres://%s:%s@%s/%s",
			postgresEnv["POSTGRES_USER"],
			postgresEnv["POSTGRES_PASSWORD"],
			postgresEnv["POSTGRES_HOST"],
			postgresEnv["POSTGRES_DB"],
		),
		DeploymentMode:       mode,
		PublicRoutes:         allowedPublicRoutes,
		LocalPath:            storageEnvs["LOCAL_PATH"],
		MinioEndpoint:        storageEnvs["MINIO_ENDPOINT"],
		MinioBucketName:      storageEnvs["MINIO_BUCKET_NAME"],
		MinioBucketRegion:    storageEnvs["MINIO_BUCKET_REGION"],
		MinioAccessKeyId:     storageEnvs["MINIO_ACCESS_KEY_ID"],
		MinioSecretAccessKey: storageEnvs["MINIO_SECRET_ACCESS_KEY"],
		MinioSSL:             storageEnvs["MINIO_SSL"] == "true",
	}, nil
}
