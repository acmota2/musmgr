package pieces

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/policies"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeletePiece(fh *controller.FilesHandler, bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		group := "DeletePiece"
		logger := bh.Logger.WithGroup(group)

		pieceID, err := uuid.Parse(c.Param("piece_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx := c.Request.Context()

		var files []model.MusmgrFile
		err = bh.DBTransaction(ctx, func(qtx *model.Queries) error {
			queryArgs := model.GetPieceFilesParams{
				PieceID:        pieceID,
				Classification: int16(policies.MaxClassification),
			}
			files, err = qtx.GetPieceFiles(ctx, queryArgs)
			if err != nil {
				return err
			}

			return qtx.DeletePiece(ctx, pieceID)
		})
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		fileIDs := make([]uuid.UUID, 0, len(files))
		for _, file := range files {
			fileIDs = append(fileIDs, file.ID)
		}
		fh.BestEffortDelete(logger, group, fileIDs...)

		c.Status(http.StatusNoContent)
		logger.Info("success")
	}
}
