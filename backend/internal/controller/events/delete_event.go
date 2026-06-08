package events

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteEvent(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("DeleteEvent")

		eventId, err := uuid.Parse(c.Param("event_id"))
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		if err := bh.Queries.DeleteEvent(ctx, eventId); err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusNoContent)
		logger.Info("Success")
	}
}
