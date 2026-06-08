package events

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func GetEvents(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetEvents")

		ctx := c.Request.Context()
		events, err := bh.Queries.GetEvents(ctx)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, events)
		logger.Info("Success")
	}
}
