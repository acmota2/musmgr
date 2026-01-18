package platform

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"
)

const (
	MINIO = "MINIO"
	LOCAL = "LOCAL"
)

type StorageType string

type StorageConfig struct {
	LocalPath      string
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioBucket    string
}

type StorageManager interface {
	Create(ctx context.Context, id uuid.UUID, r io.Reader, size int64) error // unknown size is -1
	Read(ctx context.Context, id uuid.UUID) (io.ReadCloser, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewStorage(kind StorageType, cfg StorageConfig) (StorageManager, error) {
	switch kind {
	case LOCAL:
		return newLocalStorage(cfg.LocalPath)
	case MINIO:
	default:
	}
	return nil, fmt.Errorf("Wrong storage type: %s", kind)
}
