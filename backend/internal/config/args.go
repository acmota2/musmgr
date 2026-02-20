package config

import (
	"flag"
	"fmt"

	platform "github.com/acmota2/musmgr/backend/internal/platform/file-access"
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

func ParseStorageType(s string) (platform.StorageType, error) {
	if s == "" {
		s = platform.MINIO
	}

	switch s {
	case platform.LOCAL, platform.MINIO:
		return platform.StorageType(s), nil
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
