package controller

import (
	"github.com/acmota2/musmgr/backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateEvent(c *gin.Context) {
	var event model.CreateEventParams
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	if err := h.Queries.CreateEvent(ctx, event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": event})
	}
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	eventId, err := uuid.Parse(c.Param("event_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	if err := h.Queries.DeleteEvent(ctx, eventId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "event deleted"})
	}
}

func (h *Handler) GetEvents(c *gin.Context) {
	ctx := c.Request.Context()
	if events, err := h.Queries.GetEvents(ctx); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": events})
	}
}

func (h *Handler) GetEventPieces(c *gin.Context) {
	eventId, err := uuid.Parse(c.Param("event_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	if works, err := h.Queries.GetEventPieces(ctx, eventId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": works})
	}
}
