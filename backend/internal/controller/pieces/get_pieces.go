package pieces

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func GetPieces(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetPieces")

		ctx := c.Request.Context()

		pieces, err := bh.Queries.GetPieces(ctx)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, pieces)
		logger.Info("success")
	}
}
