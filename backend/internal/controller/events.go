package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/jackc/pgx/v5"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func formatEventDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

type createEventRequest struct {
	Name        string                `json:"name"`
	Description *string               `json:"description"`
	EventType   model.MusmgrEventType `json:"event_type"`
	HappenedAt  string                `json:"happened_at"`
}

func (h *Handler) CreateEvent(c *gin.Context) {
	logger := h.Logger.WithGroup("CreateEvent")

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
		Description: textOrNull(req.Description),
		HappenedAt:  req.HappenedAt,
		EventType:   req.EventType,
	}

	if err := h.Queries.CreateEvent(ctx, queryArgs); err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Location", fmt.Sprintf("/events/%s", newEventID))
	c.Status(http.StatusCreated)
	logger.Info("Event created", "id", newEventID)
}

type updateEventRequest struct {
	Name        *string                `json:"name"`
	Description *string                `json:"description"`
	EventType   *model.MusmgrEventType `json:"event_type"`
	HappenedAt  *string                `json:"happened_at"`
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	logger := h.Logger.WithGroup("UpdateEvent")

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
		Name:        textOrNull(req.Name),
		Description: textOrNull(req.Description),
		HappenedAt:  textOrNull(req.HappenedAt),
	}
	if req.EventType != nil {
		queryArgs.EventType = model.NullMusmgrEventType{
			MusmgrEventType: *req.EventType,
			Valid:           true,
		}
	}

	ctx := c.Request.Context()

	err = h.Queries.UpdateEvent(ctx, queryArgs)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
	logger.Info("Success")
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	logger := h.Logger.WithGroup("DeleteEvent")

	eventId, err := uuid.Parse(c.Param("event_id"))
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	if err := h.Queries.DeleteEvent(ctx, eventId); err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
	logger.Info("Success")
}

func (h *Handler) GetEvent(c *gin.Context) {
	logger := h.Logger.WithGroup("GetEvent")

	eventID, err := uuid.Parse(c.Param("event_id"))
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	event, err := h.Queries.GetEvent(ctx, eventID)
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

func (h *Handler) GetEvents(c *gin.Context) {
	logger := h.Logger.WithGroup("GetEvents")

	ctx := c.Request.Context()
	events, err := h.Queries.GetEvents(ctx)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, events)
	logger.Info("Success")
}

func (h *Handler) GetEventPieces(c *gin.Context) {
	logger := h.Logger.WithGroup("GetEventPieces")

	eventId, err := uuid.Parse(c.Param("event_id"))
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	pieces, err := h.Queries.GetEventPieces(ctx, eventId)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, pieces)
	logger.Info("Success")
}
