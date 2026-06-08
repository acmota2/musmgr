package main

import (
	"context"
	"log"
	"os"

	fixture "github.com/acmota2/musmgr/backend/cmd/populate/fixtures"
	"github.com/acmota2/musmgr/backend/internal/config"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/platform/storage"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	storageType, err := config.ParseStorageType(os.Getenv("STORAGE_TYPE"))

	envConfig, err := config.LoadFromEnv(storageType)
	if err != nil {
		log.Panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, envConfig.DatabaseUrl)
	if err != nil {
		log.Panic(err)
	}
	defer pool.Close()

	storageCfg := storage.StorageConfig{
		Kind:                 storageType,
		LocalPath:            envConfig.LocalPath,
		MinioEndpoint:        envConfig.MinioEndpoint,
		MinioAccessKeyId:     envConfig.MinioAccessKeyId,
		MinioSecretAccessKey: envConfig.MinioSecretAccessKey,
		MinioBucketName:      envConfig.MinioBucketName,
		MinioBucketRegion:    envConfig.MinioBucketRegion,
		MinioSSL:             envConfig.MinioSSL,
	}
	strg, err := storage.NewStorage(&storageCfg)
	if err != nil {
		log.Panic(err)
	}

	queries := model.New(pool)

	if err = queries.CreateComposer(ctx, model.CreateComposerParams{
		FullName:  "John Doe",
		Biography: "My name is John Doe and I studied composition. The End.",
	}); err != nil {
		log.Panic(err)
	}

	for _, e := range fixture.MusmgrEvents {
		if err = queries.CreateEvent(ctx, e); err != nil {
			log.Panic(err)
		}
	}

	for _, p := range fixture.Pieces {
		if err = queries.CreatePiece(ctx, p); err != nil {
			log.Panic(err)
		}
	}

	for _, e := range fixture.MusmgrEvents {
		if err = queries.CreatePieceEvent(ctx, model.CreatePieceEventParams{
			PieceID: fixture.Pieces[0].ID,
			EventID: e.ID,
		}); err != nil {
			log.Panic(err)
		}
	}
	if err = queries.CreatePieceEvent(ctx, model.CreatePieceEventParams{
		PieceID: fixture.Pieces[1].ID,
		EventID: fixture.MusmgrEvents[1].ID,
	}); err != nil {
		log.Panic(err)
	}

	for path, fm := range fixture.Files {
		if err = queries.CreateFile(ctx, fm); err != nil {
			log.Panic(err)
		}

		rd, err := os.Open(path)
		if err != nil {
			log.Panic(err)
		}

		if err = strg.Create(ctx, fm.ID, rd, storage.UnknownSize, fm.ContentType); err != nil {
			log.Panic(err)
		}
	}
}
