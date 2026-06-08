package composer

import (
	"io"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func GetComposerPicture(fh *controller.FilesHandler, bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetComposerPicture")
		ctx := c.Request.Context()

		composer, err := bh.Queries.GetComposer(ctx)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if !composer.Picture.Valid {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		rd, err := fh.Storage.Read(ctx, composer.Picture.Bytes)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Header("Content-Type", "application/octet-stream")
		_, err = io.Copy(c.Writer, rd)
		if err != nil {
			logger.Info("Client disconnected, ignoring")
		}

		c.Status(http.StatusOK)
		logger.Info("success")
	}
}
