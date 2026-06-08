package enumscontroller

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func GetFileTypes(h *controller.BaseHandler) gin.HandlerFunc {
	logger := h.Logger
	queries := h.Queries

	return func(c *gin.Context) {
		logger := logger.WithGroup("GetFileTypes")

		ctx := c.Request.Context()

		in, err := queries.GetFileTypes(ctx)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, in)
		logger.Info("success")
	}
}
