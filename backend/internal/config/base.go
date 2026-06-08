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
			Kind:                 parsedArgs.StorageType,
			LocalPath:            environmentVariables.LocalPath,
			MinioEndpoint:        environmentVariables.MinioEndpoint,
			MinioAccessKeyId:     environmentVariables.MinioAccessKeyId,
			MinioSecretAccessKey: environmentVariables.MinioSecretAccessKey,
			MinioBucketName:      environmentVariables.MinioBucketName,
			MinioBucketRegion:    environmentVariables.MinioBucketRegion,
			MinioSSL:             environmentVariables.MinioSSL,
		},
	}, nil
}
