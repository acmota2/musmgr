package composer

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func UpdateComposerPicture(fh *controller.FilesHandler, bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		group := "UpdateComposerPicture"
		logger := bh.Logger.WithGroup(group)

		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": "'file' is required"},
			)
			return
		}

		// change this to another place when accepting more images
		const maxPictureSize = 2 << 20 // 2MiB
		if fileHeader.Size > maxPictureSize {
			c.AbortWithStatusJSON(
				http.StatusRequestEntityTooLarge,
				gin.H{"error": "Maximum image size is 2MiB"},
			)
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer file.Close()

		ctx := c.Request.Context()

		contentType := fileHeader.Header.Get("Content-Type")
		switch contentType {
		case "image/jpeg", "image/png", "image/webp":
		default:
			c.AbortWithStatus(http.StatusUnsupportedMediaType)
			return
		}

		err = bh.DBTransaction(ctx, func(qtx *model.Queries) error {
			composer, err := qtx.GetComposer(ctx)
			if err != nil {
				return err
			}

			newPictureID := uuid.New()
			if err = fh.Storage.Create(ctx, newPictureID, file, fileHeader.Size, contentType); err != nil {
				return err
			}

			queryArgs := model.UpdateComposerPictureParams{
				Picture:            pgtype.UUID{Bytes: newPictureID, Valid: true},
				PictureContentType: pgtype.Text{String: contentType, Valid: true},
			}
			if err = qtx.UpdateComposerPicture(ctx, queryArgs); err != nil {
				// try to delete
				fh.BestEffortDelete(logger, group, newPictureID)
				return err
			}
			if composer.Picture.Valid {
				fh.BestEffortDelete(logger, group, composer.Picture.Bytes)
			}

			return nil
		})
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Header("Location", "/composer/picture")
		c.Status(http.StatusNoContent)
		logger.Info("success")
	}
}
