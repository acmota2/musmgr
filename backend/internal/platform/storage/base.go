package storage

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/google/uuid"
)

type StorageType string

const (
	S3    = "S3"
	LOCAL = "LOCAL"
)

type StorageConfig struct {
	Kind              StorageType
	LocalPath         string
	S3Endpoint        string
	S3AccessKeyId     string
	S3SecretAccessKey string
	S3BucketName      string
	S3BucketRegion    string
	S3SSL             bool
}

const UnknownSize int64 = -1

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
	case S3:
		log.Printf("Using S3-compatible storage at: %s", cfg.S3Endpoint)
		return newS3Storage(cfg)
	default:
	}
	return nil, fmt.Errorf("wrong storage type: %s", cfg.Kind)
}
