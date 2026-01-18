package controller

import (
	"github.com/acmota2/musmgr/backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreatePiece(c *gin.Context) {
	var piece model.CreatePieceParams
	if err := c.ShouldBindJSON(&piece); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	if err := h.Queries.CreatePiece(ctx, piece); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": piece})
	}
}

func (h *Handler) DeletePiece(c *gin.Context) {
	pieceId, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	if err := h.Queries.DeletePiece(ctx, pieceId); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "piece deleted"})
	}
}

func (h *Handler) GetPieces(c *gin.Context) {
	ctx := c.Request.Context()
	if pieces, err := h.Queries.GetPieces(ctx); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": pieces})
	}
}

func (h *Handler) GetPieceEvents(c *gin.Context) {
	id, err := uuid.Parse(c.Param("piece_id"))
	if err == nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	if events, err := h.Queries.GetPieceEvents(ctx, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": events})
	}
}
