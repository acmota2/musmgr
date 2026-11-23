package filescontroller

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/middleware"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/platform/storage"
	"github.com/gin-gonic/gin"
)

func GetFile(fh *controller.FilesHandler, bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetFile")

		file, ok := middleware.GetContextValue[model.MusmgrFile](c, middleware.CurrentFile)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx := c.Request.Context()

		url, err := fh.Storage.GetURL(ctx, file.ID, storage.DefaultPresignExpiration)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Redirect(http.StatusFound, url)
	}
}
