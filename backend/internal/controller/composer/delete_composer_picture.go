package composer

import (
	"errors"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func DeleteComposerPicture(fh *controller.FilesHandler, bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		group := "DeleteComposerPicture"
		logger := bh.Logger.WithGroup(group)

		ctx := c.Request.Context()
		pictureID, err := bh.Queries.DeleteComposerPicture(ctx)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				c.Status(http.StatusNoContent)
				logger.Info("success: picture didn't exist")
				return
			}
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if pictureID.Valid {
			fh.BestEffortDelete(logger, group, pictureID.Bytes)
		}

		c.Status(http.StatusNoContent)
		logger.Info("success")
	}
}
