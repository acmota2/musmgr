package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func formatEventDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

type createEventRequest struct {
	Description *string               `json:"description"`
	EventType   model.MusmgrEventType `json:"event_type"`
	HappenedAt  string                `json:"happened_at"`
}

func (h *Handler) CreateEvent(c *gin.Context) {
	var req createEventRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	happenedAt, err := formatEventDate(req.HappenedAt)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	newEventID := uuid.New()
	queryArgs := model.CreateEventParams{
		ID:          newEventID,
		Description: textOrNull(req.Description),
		HappenedAt: pgtype.Date{
			Time:  happenedAt,
			Valid: true,
		},
		EventType: req.EventType,
	}

	if err := h.Queries.CreateEvent(ctx, queryArgs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Location", fmt.Sprintf("/events/%s", newEventID))
	c.Status(http.StatusCreated)
}

type updateEventRequest struct {
	Description *string                `json:"description"`
	EventType   *model.MusmgrEventType `json:"event_type"`
	HappenedAt  *string                `json:"happened_at"`
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	eventID, err := uuid.Parse(c.Param("event_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req updateEventRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	queryArgs := model.UpdateEventParams{
		ID:          eventID,
		Description: textOrNull(req.Description),
	}

	if req.HappenedAt != nil {
		happenedAt, err := formatEventDate(*req.HappenedAt)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		queryArgs.HappenedAt = pgtype.Date{
			Time:  happenedAt,
			Valid: true,
		}
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
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	eventId, err := uuid.Parse(c.Param("event_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	if err := h.Queries.DeleteEvent(ctx, eventId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) GetEvent(c *gin.Context) {
	eventID, err := uuid.Parse(c.Param("event_id"))
	if err != nil {
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
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *Handler) GetEvents(c *gin.Context) {
	ctx := c.Request.Context()
	events, err := h.Queries.GetEvents(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, events)
}

func (h *Handler) GetEventPieces(c *gin.Context) {
	eventId, err := uuid.Parse(c.Param("event_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	pieces, err := h.Queries.GetEventPieces(ctx, eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pieces)
}
