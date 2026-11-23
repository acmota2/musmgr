package composer

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/platform/storage"
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

		url, err := fh.Storage.GetURL(ctx, composer.Picture.Bytes, storage.DefaultPresignExpiration)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Redirect(http.StatusFound, url)
		logger.Info("success")
	}
}
