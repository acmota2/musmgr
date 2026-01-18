package platform

import "fmt"

const (
	MINIO = "MINIO"
	LOCAL = "LOCAL"
)

type StorageType string

type StorageManager interface {
	New(string)
	Update(string)
	Delete(string)
}

func NewStorage(kind StorageType) {
	switch kind {
	case LOCAL:
	case MINIO:
	default:
	}
}

func ParseStorageType(s string) (StorageType, error) {
	if s == "" {
		s = "MINIO"
	}

	switch s {
	case "MINIO", "LOCAL":
		return StorageType(s), nil
	default:
		return "", fmt.Errorf("Invalid storage type: %q", s)
	}
}
