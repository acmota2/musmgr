package controller

import (
	"github.com/acmota2/musmgr/backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePieceEvent(c *gin.Context) {
	var workEvent model.CreatePieceEventParams
	if err := c.ShouldBindJSON(&workEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	if err := h.Queries.CreatePieceEvent(ctx, workEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": workEvent})
	}
}
