package events

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type updateEventRequest struct {
	Name        *string                `json:"name"`
	Description *string                `json:"description"`
	EventType   *model.MusmgrEventType `json:"event_type"`
	HappenedAt  *string                `json:"happened_at"`
}

func UpdateEvent(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("UpdateEvent")

		eventID, err := uuid.Parse(c.Param("event_id"))
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var req updateEventRequest
		if err = c.ShouldBindJSON(&req); err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		_, err = formatEventDate(*req.HappenedAt)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		queryArgs := model.UpdateEventParams{
			ID:          eventID,
			Name:        controller.TextOrNull(req.Name),
			Description: controller.TextOrNull(req.Description),
			HappenedAt:  controller.TextOrNull(req.HappenedAt),
		}
		if req.EventType != nil {
			queryArgs.EventType = model.NullMusmgrEventType{
				MusmgrEventType: *req.EventType,
				Valid:           true,
			}
		}

		ctx := c.Request.Context()

		err = bh.Queries.UpdateEvent(ctx, queryArgs)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusNoContent)
		logger.Info("Success")
	}
}
