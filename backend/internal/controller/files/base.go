package filescontroller

import (
	"context"
	"log/slog"
	"time"

	"github.com/acmota2/musmgr/backend/internal/platform/storage"
	services "github.com/acmota2/musmgr/backend/internal/services/pdf-generation"
	"github.com/google/uuid"
)

// FilesHandler wraps storage and PDF generation dependencies for file operations.
type FilesHandler struct {
	Storage      storage.StorageManager
	PdfGenerator services.PdfGenerator
}

// this will be a queue effort in the future
func (fh *FilesHandler) BestEffortDelete(logger *slog.Logger, group string, ids ...uuid.UUID) {
	go func() {
		logger = logger.WithGroup(group)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		for _, id := range ids {
			err := fh.Storage.Delete(ctx, id)
			if err != nil {
				logger.Warn("Tried to delete file", "id", id, "err", err)
			}
		}
	}()
}
