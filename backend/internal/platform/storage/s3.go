package storage

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type s3Storage struct {
	client     *s3.Client
	bucketName string
}

func newS3Storage(sc *StorageConfig) (*s3Storage, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(sc.S3BucketRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			sc.S3AccessKeyId,
			sc.S3SecretAccessKey,
			"",
		)),
	)
	if err != nil {
		return nil, err
	}

	endpoint := "http://" + sc.S3Endpoint
	if sc.S3SSL {
		endpoint = "https://" + sc.S3Endpoint
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(sc.S3BucketName),
	})
	if err != nil {
		_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(sc.S3BucketName),
		})
		if err != nil {
			return nil, err
		}
	}

	return &s3Storage{
		client:     client,
		bucketName: sc.S3BucketName,
	}, nil
}

func (s *s3Storage) Create(ctx context.Context, id uuid.UUID, r io.Reader, size int64, contentType string) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(id.String()),
		Body:        r,
		ContentType: aws.String(contentType),
	})
	return err
}

func (s *s3Storage) Read(ctx context.Context, id uuid.UUID) (io.ReadCloser, error) {
	result, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(id.String()),
	})
	if err != nil {
		return nil, err
	}
	return result.Body, nil
}

func (s *s3Storage) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(id.String()),
	})
	return err
}
