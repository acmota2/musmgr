package composer

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateComposer(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("CreateComposer")

		var req model.CreateComposerParams
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		err := bh.Queries.CreateComposer(ctx, req)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Header("Location", "/composer")
		c.Status(http.StatusCreated)
		logger.Info("success")
	}
}
