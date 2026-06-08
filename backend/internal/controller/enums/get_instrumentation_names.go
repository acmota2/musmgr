package enumscontroller

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func GetInstrumentationNames(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("GetInstrumentationNames")

		ctx := c.Request.Context()

		in, err := bh.Queries.GetInstrumentationNames(ctx)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, in)
		logger.Info("success")
	}
}
