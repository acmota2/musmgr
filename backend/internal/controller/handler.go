package controller

import (
	"context"
	"log/slog"
	"time"

	"github.com/acmota2/musmgr/backend/internal/model"
	platform "github.com/acmota2/musmgr/backend/internal/platform/file-access"
	services "github.com/acmota2/musmgr/backend/internal/services/pdf-generation"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Pool         *pgxpool.Pool
	Queries      *model.Queries
	Storage      platform.StorageManager
	PdfGenerator services.PdfGenerator
	Logger       *slog.Logger
}

func (h *Handler) BestEffortDelete(group string, ids ...uuid.UUID) {
	go func() {
		logger := h.Logger.WithGroup(group)

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

func (h *Handler) DBTransaction(ctx context.Context, f func(qtx *model.Queries) error) error {
	tx, err := h.Pool.Begin(ctx)
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
