package filescontroller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/platform/storage"
	"github.com/acmota2/musmgr/backend/internal/policies"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type createFileRequest struct {
	Classification string `form:"classification" binding:"required"`
	FileType       string `form:"file_type" binding:"required"`
	Name           string `form:"name" binding:"required"`
}

func createPreviewScore(ctx context.Context, fh *controller.FilesHandler, bh *controller.BaseHandler, id uuid.UUID, fileName string, pieceID uuid.UUID) (*uuid.UUID, error) {
	rd, err := fh.Storage.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rd.Close()

	preview, err := fh.PdfGenerator.Generate(ctx, rd)
	if err != nil {
		return nil, err
	}
	defer preview.Close()

	newID := uuid.New()

	err = fh.Storage.Create(ctx, newID, preview, storage.UnknownSize, "application/pdf")
	if err != nil {
		return nil, err
	}

	queryParams := model.CreateFileParams{
		ID:             newID,
		Classification: policies.ScopePublic,
		ContentType:    "application/pdf",
		FileType:       model.MusmgrFileTypeScorePreview,
		Name:           fileName,
		Origin:         model.MusmgrFileOriginSystem,
		ParentID:       pgtype.UUID{Valid: true, Bytes: id},
		PieceID:        pieceID,
	}

	if err = bh.Queries.CreateFile(ctx, queryParams); err != nil {
		return nil, err
	}

	return &newID, nil
}

func CreateFile(fh *controller.FilesHandler, bh *controller.BaseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := bh.Logger.WithGroup("CreateFile")

		pieceID, err := uuid.Parse(c.Param("piece_id"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var fileParams createFileRequest
		if err = c.ShouldBindWith(&fileParams, binding.FormMultipart); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		class, err := policies.StringToClassification(fileParams.Classification)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		fileType := model.MusmgrFileType(fileParams.FileType)
		if !fileType.Valid() {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		newFileID := uuid.New()
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
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
		if contentType == "" {
			contentType = "application/octet-stream"
		}

		queryParams := model.CreateFileParams{
			ID:             newFileID,
			Classification: int16(class),
			ContentType:    contentType,
			FileType:       fileType,
			Name:           fileParams.Name,
			Origin:         model.MusmgrFileOriginUser,
			ParentID:       pgtype.UUID{Valid: false},
			PieceID:        pieceID,
		}

		err = bh.Queries.CreateFile(ctx, queryParams)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		err = fh.Storage.Create(ctx, newFileID, file, fileHeader.Size, contentType)
		if err != nil {
			logger.Error(err.Error())

			delErr := bh.Queries.DeleteFile(ctx, newFileID)
			if delErr != nil {
				logger.Warn("couldn't delete file", "err", delErr)
			}

			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Header("Location", fmt.Sprintf("/pieces/%s/files/%s", pieceID, newFileID))
		c.JSON(http.StatusCreated, gin.H{
			"id": newFileID,
		})
		logger.Info("success creating file")

		if fileType == model.MusmgrFileTypeScoreFull {
			previewID, err := createPreviewScore(ctx, fh, bh, newFileID, fileParams.Name, pieceID)
			if err != nil {
				logger.Error("didn't create preview file", "id", newFileID, "err", err)
				return
			}
			logger.Info("created preview", "id", newFileID, "preview_id", *previewID)
		}
	}
}
