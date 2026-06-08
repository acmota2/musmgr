package filescontroller

import (
	"errors"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/middleware"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func DeleteFile(fh *controller.FilesHandler, bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("DeleteFile")

		file, ok := middleware.GetContextValue[model.MusmgrFile](c, middleware.CurrentFile)
		if !ok {
			logger.Error("While converting file type")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx := c.Request.Context()

		err := bh.Queries.DeleteFile(ctx, file.ID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				logger.Info("File doesn't exist")
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = fh.Storage.Delete(ctx, file.ID)
		if err != nil {
			logger.Warn("Failed to delete file", "id", file.ID, "err", err)
		}

		c.Status(http.StatusNoContent)
		logger.Info("Success")
	}
}
