package pieces

import (
	"fmt"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createPieceRequest struct {
	ComposedAt      string                          `json:"composed_at"`
	Description     string                          `json:"description"`
	Instrumentation model.MusmgrInstrumentationName `json:"instrumentation"`
	Title           string                          `json:"title"`
}

func CreatePiece(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("CreatePiece")

		var req createPieceRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		_, err := formatPieceDate(req.ComposedAt)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		newID := uuid.New()
		queryArgs := model.CreatePieceParams{
			ID:              newID,
			ComposedAt:      req.ComposedAt,
			Description:     req.Description,
			Instrumentation: req.Instrumentation,
			Title:           req.Title,
		}

		if err := bh.Queries.CreatePiece(ctx, queryArgs); err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Header("Location", fmt.Sprintf("/pieces/%s", newID.String()))
		c.Status(http.StatusCreated)
		logger.Info("success")
	}
}
