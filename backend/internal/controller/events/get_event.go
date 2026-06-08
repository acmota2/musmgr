package events

import (
	"errors"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func GetEvent(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetEvent")

		eventID, err := uuid.Parse(c.Param("event_id"))
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx := c.Request.Context()

		event, err := bh.Queries.GetEvent(ctx, eventID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, event)
		logger.Info("Success")
	}
}
