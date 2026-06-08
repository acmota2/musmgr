package controller

import (
	"context"
	"log/slog"
	"time"

	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/platform/storage"
	services "github.com/acmota2/musmgr/backend/internal/services/pdf-generation"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BaseHandler struct {
	pool    *pgxpool.Pool
	Queries *model.Queries
	Logger  *slog.Logger
}

func NewBaseHandler(pool *pgxpool.Pool, queries *model.Queries, logger *slog.Logger) *BaseHandler {
	return &BaseHandler{
		pool:    pool,
		Queries: queries,
		Logger:  logger,
	}
}

func (h *BaseHandler) DBTransaction(ctx context.Context, f func(qtx *model.Queries) error) error {
	tx, err := h.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := model.New(tx)
	if err = f(qtx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

type FilesHandler struct {
	Storage      storage.StorageManager
	PdfGenerator services.PdfGenerator
}

// this will be a queue effort in the future
func (h *FilesHandler) BestEffortDelete(logger *slog.Logger, group string, ids ...uuid.UUID) {
	go func() {
		logger = logger.WithGroup(group)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		for _, id := range ids {
			err := h.Storage.Delete(ctx, id)
			if err != nil {
				logger.Warn("Tried to delete file", "id", id, "err", err)
			}
		}
	}()
}
