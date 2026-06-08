package controller

import (
	"context"
	"log/slog"

	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BaseHandler struct {
	pool    *pgxpool.Pool
	Queries *model.Queries
	Logger  *slog.Logger
}

func New(pool *pgxpool.Pool, queries *model.Queries, logger *slog.Logger) *BaseHandler {
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
