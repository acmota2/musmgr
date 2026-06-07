package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetEventTypes(c *gin.Context) {
	logger := h.Logger.WithGroup("GetEventTypes")

	ctx := c.Request.Context()

	in, err := h.Queries.GetEventTypes(ctx)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, in)
	logger.Info("success")
}
