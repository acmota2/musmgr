package config

import (
	platform "backend/internal/platform/file_access"
	"flag"
	"log"
)

type ArgsError struct {
	Message string
}

type fromArgsConfig struct {
	EnvFilePath string
	StorageType platform.StorageType
	AdminPort   string
	PublicPort  string
}

func loadFromArgs() (fromArgsConfig, error) {
	var argsConfig fromArgsConfig
	var storageType string
	flag.StringVar(&argsConfig.EnvFilePath, "env-file", "", "Path to the .env file")
	flag.StringVar(&storageType, "storage-type", "MINIO", "Type of storage to use: MINIO | LOCAL")
	flag.StringVar(&argsConfig.AdminPort, "admin-port", "4700", "The port where the admin backend should run")
	flag.StringVar(&argsConfig.PublicPort, "public-port", "4701", "The port where the public backend should run")
	flag.Parse()

	newStorageType, err := platform.ParseStorageType(storageType)
	if err != nil {
		newStorageType = platform.MINIO
		log.Println("Invalid storage type, will use MinIO")
	}

	argsConfig.StorageType = newStorageType

	return argsConfig, nil
}
