package events

import (
	"fmt"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createEventRequest struct {
	Name        string                `json:"name"`
	Description *string               `json:"description"`
	EventType   model.MusmgrEventType `json:"event_type"`
	HappenedAt  string                `json:"happened_at"`
}

func CreateEvent(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("CreateEvent")

		var req createEventRequest
		if err := c.BindJSON(&req); err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		_, err := formatEventDate(req.HappenedAt)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		newEventID := uuid.New()
		queryArgs := model.CreateEventParams{
			ID:          newEventID,
			Name:        req.Name,
			Description: controller.TextOrNull(req.Description),
			HappenedAt:  req.HappenedAt,
			EventType:   req.EventType,
		}

		if err := bh.Queries.CreateEvent(ctx, queryArgs); err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Header("Location", fmt.Sprintf("/events/%s", newEventID))
		c.Status(http.StatusCreated)
		logger.Info("Event created", "id", newEventID)
	}
}
