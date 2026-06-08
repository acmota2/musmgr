package filescontroller

import (
	"io"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/middleware"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
)

// GetFile returns a Gin handler that streams a file's content.
func GetFile(fh *FilesHandler, bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetFile")

		file, ok := middleware.GetContextValue[model.MusmgrFile](c, middleware.CurrentFile)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx := c.Request.Context()

		rd, err := fh.Storage.Read(ctx, file.ID)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer rd.Close()

		c.Header("Content-Type", file.ContentType)
		c.Status(http.StatusOK)

		_, err = io.Copy(c.Writer, rd)
		if err != nil {
			logger.Info("Client disconnected")
			return
		}
	}
}
