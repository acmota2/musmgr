package controller

import (
	"github.com/acmota2/musmgr/backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateWork(c *gin.Context) {
	var work model.CreateWorkParams
	if err := c.ShouldBindJSON(&work); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	if err := h.Queries.CreateWork(ctx, work); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": work})
}

func (h *Handler) DeleteWork(c *gin.Context) {
	id, err := uuid.Parse(c.Param("work_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	if err := h.Queries.DeleteWork(ctx, id.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "work deleted"})
	}
}

func (h *Handler) GetWorks(c *gin.Context) {
	ctx := c.Request.Context()
	if works, err := h.Queries.GetWorks(ctx); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": works})
	}
}

func (h *Handler) GetWorkEvents(c *gin.Context) {
	id := c.Param("work_id")
	if _, err := uuid.Parse(id); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid work_id"})
		return
	}

	ctx := c.Request.Context()

	if events, err := h.Queries.GetWorkEvents(ctx, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": events})
	}
}
