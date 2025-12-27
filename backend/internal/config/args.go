package config

import (
	"flag"
)

type StorageType string
const (
	Minio StorageType = "MINIO"
	Local StorageType = "LOCAL"
)

type DeploymentType string
const (
	Development DeploymentType = "DEVELOPMENT"
	Production  DeploymentType = "PRODUCTION"
)

type fromArgsConfig struct {
	DeploymentType	DeploymentType
	EnvFilePath string
	StorageType StorageType
}

func loadFromArgs() fromArgsConfig {
	var argsConfig fromArgsConfig
	var storageType, deploymentType string
	flag.StringVar(&argsConfig.EnvFilePath, "env-file", "", "Path to the .env file")
	flag.StringVar(&storageType, "storage-type", "LOCAL", "Type of storage to use: MINIO | LOCAL")
	flag.StringVar(&deploymentType, "storage", "DEVELOPMENT", "Deployment type: DEVELOPMENT | PRODUCTION")
	flag.Parse()

	argsConfig.DeploymentType = DeploymentType(deploymentType)
  argsConfig.StorageType = StorageType(storageType)

	return argsConfig
}
