package filescontroller

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type updateFileMetadataRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func UpdateFileMetadata(bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("UpdateFileMetadata")

		fileID, err := uuid.Parse(c.Param("file_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		var req updateFileMetadataRequest
		if err = c.ShouldBindJSON(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		queryArgs := model.UpdateFileMetadataParams{
			ID:          fileID,
			Description: controller.TextOrNull(req.Description),
			Name:        controller.TextOrNull(req.Name),
		}

		if err = bh.Queries.UpdateFileMetadata(ctx, queryArgs); err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusNoContent)
		logger.Info("updated")
	}
}
