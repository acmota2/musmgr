package config

import (
	"flag"
	"fmt"

	"github.com/acmota2/musmgr/backend/internal/platform/storage"
)

type ArgsError struct {
	Message string
}

type configFromArgs struct {
	EnvFilePath string
	StorageType storage.StorageType
	AdminPort   string
	PublicPort  string
}

func ParseStorageType(s string) (storage.StorageType, error) {
	if s == "" {
		s = storage.MINIO
	}

	switch s {
	case storage.LOCAL, storage.MINIO:
		return storage.StorageType(s), nil
	default:
		return "", fmt.Errorf("Invalid storage type: %s", s)
	}
}

func loadFromArgs() (configFromArgs, error) {
	var argsConfig configFromArgs
	var storageType string
	flag.StringVar(&storageType, "storage-type", "MINIO", "Type of storage to use: MINIO | LOCAL")
	flag.StringVar(&argsConfig.AdminPort, "admin-port", "4700", "The port where the admin backend should run")
	flag.StringVar(&argsConfig.PublicPort, "public-port", "4701", "The port where the public backend should run")
	flag.Parse()

	newStorageType, err := ParseStorageType(storageType)
	if err != nil {
		return configFromArgs{}, err
	}

	argsConfig.StorageType = newStorageType

	return argsConfig, nil
}
