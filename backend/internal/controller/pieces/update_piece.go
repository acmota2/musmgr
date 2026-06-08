package pieces

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type updatePieceRequest struct {
	ComposedAt      *string                          `json:"composed_at"`
	Description     *string                          `json:"description"`
	Instrumentation *model.MusmgrInstrumentationName `json:"instrumentation"`
	Title           *string                          `json:"title"`
}

func UpdatePiece(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("UpdatePiece")

		pieceID, err := uuid.Parse(c.Param("piece_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var req updatePieceRequest
		if err = c.ShouldBindJSON(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		_, err = formatPieceDate(*req.ComposedAt)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		queryArgs := model.UpdatePieceParams{
			ID:          pieceID,
			Description: controller.TextOrNull(req.Description),
			Title:       controller.TextOrNull(req.Title),
			ComposedAt:  controller.TextOrNull(req.ComposedAt),
		}

		if req.Instrumentation != nil {
			queryArgs.Instrumentation = model.NullMusmgrInstrumentationName{
				MusmgrInstrumentationName: *req.Instrumentation,
				Valid:                     true,
			}
		}

		err = bh.Queries.UpdatePiece(ctx, queryArgs)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusNoContent)
		logger.Info("success")
	}
}
