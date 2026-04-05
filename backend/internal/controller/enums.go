package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInstrumentationNames(c *gin.Context) {
	logger := h.Logger.WithGroup("GetInstrumentationNames")

	ctx := c.Request.Context()

	in, err := h.Queries.GetInstrumentationNames(ctx)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, in)
	logger.Info("success")
}

func (h *Handler) GetFileTypes(c *gin.Context) {
	logger := h.Logger.WithGroup("GetFileTypes")

	ctx := c.Request.Context()

	in, err := h.Queries.GetFileTypes(ctx)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, in)
	logger.Info("success")
}

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
