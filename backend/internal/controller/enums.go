package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInstrumentationNames(c *gin.Context) {
	ctx := c.Request.Context()

	in, err := h.Queries.GetInstrumentationNames(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, in)
}

func (h *Handler) GetFileTypes(c *gin.Context) {
	ctx := c.Request.Context()

	in, err := h.Queries.GetFileTypes(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, in)
}

func (h *Handler) GetEventTypes(c *gin.Context) {
	ctx := c.Request.Context()

	in, err := h.Queries.GetEventTypes(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, in)
}
