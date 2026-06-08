package pieces

import (
	"errors"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func GetPiece(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetPiece")

		pieceID, err := uuid.Parse(c.Param("piece_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		piece, err := bh.Queries.GetPiece(ctx, pieceID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, piece)
		logger.Info("success")
	}
}
