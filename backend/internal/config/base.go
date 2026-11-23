package config

import "github.com/acmota2/musmgr/backend/internal/platform/storage"

type Config struct {
	DatabaseUrl    string
	AdminPort      string
	PublicPort     string
	AdminRoutes    []string
	PublicRoutes   []string
	StorageConfig  storage.StorageConfig
	TrustedProxies []string
}

func New() (Config, error) {
	parsedArgs, err := loadFromArgs()
	if err != nil {
		return Config{}, err
	}

	environmentVariables, err := LoadFromEnv(parsedArgs.StorageType)
	if err != nil {
		return Config{}, err
	}

	return Config{
		DatabaseUrl:    environmentVariables.DatabaseUrl,
		AdminPort:      parsedArgs.AdminPort,
		PublicPort:     parsedArgs.PublicPort,
		AdminRoutes:    environmentVariables.AdminRoutes,
		PublicRoutes:   environmentVariables.PublicRoutes,
		TrustedProxies: environmentVariables.TrustedProxies,
		StorageConfig: storage.StorageConfig{
			Kind:              parsedArgs.StorageType,
			LocalPath:         environmentVariables.LocalPath,
			S3Endpoint:        environmentVariables.S3Endpoint,
			S3AccessKeyId:     environmentVariables.S3AccessKeyId,
			S3SecretAccessKey: environmentVariables.S3SecretAccessKey,
			S3BucketName:      environmentVariables.S3BucketName,
			S3BucketRegion:    environmentVariables.S3BucketRegion,
			S3SSL:             environmentVariables.S3SSL,
		},
	}, nil
}
