package composer

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
)

type updateComposerRequest struct {
	FullName  *string `json:"full_name"`
	Biography *string `json:"biography"`
}

func UpdateComposer(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("UpdateComposer")

		var req updateComposerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		queryArgs := model.UpdateComposerParams{
			FullName:  controller.TextOrNull(req.FullName),
			Biography: controller.TextOrNull(req.Biography),
		}
		if err := bh.Queries.UpdateComposer(ctx, queryArgs); err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Header("Location", "/composer")
		c.Status(http.StatusNoContent)
		logger.Info("success")
	}
}
