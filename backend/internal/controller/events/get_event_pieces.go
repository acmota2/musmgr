package events

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetEventPieces(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetEventPieces")

		eventId, err := uuid.Parse(c.Param("event_id"))
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		pieces, err := bh.Queries.GetEventPieces(ctx, eventId)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, pieces)
		logger.Info("Success")
	}
}
