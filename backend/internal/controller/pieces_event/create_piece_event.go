package piecesevent

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func CreatePieceEvent(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("CreatePieceEvent")

		pieceID, err := uuid.Parse(c.Param("piece_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		eventID, err := uuid.Parse(c.Param("event_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		queryArgs := model.CreatePieceEventParams{
			PieceID: pieceID,
			EventID: eventID,
		}

		if err := bh.Queries.CreatePieceEvent(ctx, queryArgs); err != nil {
			logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusCreated)
		logger.Info("success")
	}
}
