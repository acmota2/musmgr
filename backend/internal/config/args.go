package config

import (
	"flag"
	"fmt"
	"log"

	platform "github.com/acmota2/musmgr/backend/internal/platform/file_access"
)

type ArgsError struct {
	Message string
}

type configFromArgs struct {
	EnvFilePath string
	StorageType platform.StorageType
	AdminPort   string
	PublicPort  string
}

func parseStorageType(s string) (platform.StorageType, error) {
	if s == "" {
		s = "MINIO"
	}

	switch s {
	case "MINIO", "LOCAL":
		return platform.StorageType(s), nil
	default:
		return "", fmt.Errorf("Invalid storage type: %q", s)
	}
}

func loadFromArgs() (configFromArgs, error) {
	var argsConfig configFromArgs
	var storageType string
	flag.StringVar(&argsConfig.EnvFilePath, "env-file", "", "Path to the .env file")
	flag.StringVar(&storageType, "storage-type", "MINIO", "Type of storage to use: MINIO | LOCAL")
	flag.StringVar(&argsConfig.AdminPort, "admin-port", "4700", "The port where the admin backend should run")
	flag.StringVar(&argsConfig.PublicPort, "public-port", "4701", "The port where the public backend should run")
	flag.Parse()

	newStorageType, err := parseStorageType(storageType)
	if err != nil {
		newStorageType = platform.MINIO
		log.Println("Invalid storage type, will use MinIO")
	}

	argsConfig.StorageType = newStorageType

	return argsConfig, nil
}
