package platform

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/google/uuid"
)

type StorageType string

const (
	MINIO = "MINIO"
	LOCAL = "LOCAL"
)

type StorageConfig struct {
	Kind                 StorageType
	LocalPath            string
	MinioEndpoint        string
	MinioAccessKeyId     string
	MinioSecretAccessKey string
	MinioBucketName      string
	MinioBucketRegion    string
	MinioSSL             bool
}

type StorageManager interface {
	Create(ctx context.Context, id uuid.UUID, r io.Reader, size int64, contentType string) error // unknown size is -1
	Read(ctx context.Context, id uuid.UUID) (io.ReadCloser, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewStorage(cfg *StorageConfig) (StorageManager, error) {
	switch cfg.Kind {
	case LOCAL:
		log.Println("Using local storage")
		return newLocalStorage(cfg.LocalPath)
	case MINIO:
		log.Printf("Using MinIO listening on: %s", cfg.MinioEndpoint)
		return newMinioStorage(cfg)
	default:
	}
	return nil, fmt.Errorf("Wrong storage type: %s", cfg.Kind)
}
