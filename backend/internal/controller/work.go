package controller

import (
	"backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWorks(c *gin.Context) {
	ctx := c.Request.Context()
	if works, err := h.Queries.GetWorks(ctx); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": works})
	}
}

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
