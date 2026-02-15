package config

import platform "github.com/acmota2/musmgr/backend/internal/platform/file-access"

type Config struct {
	DatabaseUrl   string
	AdminPort     string
	PublicPort    string
	AdminRoutes   []string
	PublicRoutes  []string
	StorageConfig platform.StorageConfig
}

func New() (Config, error) {
	parsedArgs, err := loadFromArgs()
	if err != nil {
		return Config{}, err
	}

	environmentVariables, err := loadFromEnv(parsedArgs)
	if err != nil {
		return Config{}, err
	}

	return Config{
		DatabaseUrl:  environmentVariables.DatabaseUrl,
		AdminPort:    parsedArgs.AdminPort,
		PublicPort:   parsedArgs.PublicPort,
		AdminRoutes:  environmentVariables.AdminRoutes,
		PublicRoutes: environmentVariables.PublicRoutes,
		StorageConfig: platform.StorageConfig{
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
