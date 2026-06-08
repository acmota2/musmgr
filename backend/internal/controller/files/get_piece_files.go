package filescontroller

import (
	"errors"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/middleware"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/policies"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func GetPieceFiles(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetPieceFiles")

		pieceId, err := uuid.Parse(c.Param("piece_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		classification := c.Value(middleware.ClassKey).(policies.FileClassification)
		queryArgs := model.GetPieceFilesParams{
			PieceID:        pieceId,
			Classification: int16(classification),
		}

		files, err := bh.Queries.GetPieceFiles(ctx, queryArgs)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, files)
	}
}
