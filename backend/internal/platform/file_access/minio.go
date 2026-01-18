package platform

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioStorage struct {
	MinioClient *minio.Client
	BucketName  string
}

func newMinioStorage(sc StorageConfig) (*minioStorage, error) {
	minioClient, err := minio.New(
		sc.MinioEndpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				sc.MinioAccessKeyId,
				sc.MinioSecretAccessKey,
				"",
			),
			Secure: sc.MinioSSL,
			Region: sc.MinioBucketRegion,
		},
	)
	if err != nil {
		return nil, err
	}

	minioCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exists, err := minioClient.BucketExists(minioCtx, sc.MinioBucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("Minio: bucket %s doesn't exist", sc.MinioBucketName)
	}

	return &minioStorage{
		BucketName:  sc.MinioBucketName,
		MinioClient: minioClient,
	}, err
}

func (ms *minioStorage) Create(ctx context.Context, id uuid.UUID, r io.Reader, size int64, contentType string) error {
	_, err := ms.MinioClient.PutObject(ctx, ms.BucketName, id.String(), r, size, minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (ms *minioStorage) Read(ctx context.Context, id uuid.UUID) (io.ReadCloser, error) {
	return ms.MinioClient.GetObject(ctx, ms.BucketName, id.String(), minio.GetObjectOptions{})
}

func (ms *minioStorage) Delete(ctx context.Context, id uuid.UUID) error {
	return ms.MinioClient.RemoveObject(ctx, ms.BucketName, id.String(), minio.RemoveObjectOptions{})
}
