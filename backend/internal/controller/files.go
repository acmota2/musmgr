package controller

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/middleware"
	"github.com/acmota2/musmgr/backend/internal/model"
	platform "github.com/acmota2/musmgr/backend/internal/platform/file-access"
	"github.com/acmota2/musmgr/backend/internal/policies"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) GetPieceFiles(c *gin.Context) {
	logger := h.Logger.WithGroup("GetPieceFiles")

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

	files, err := h.Queries.GetPieceFiles(ctx, queryArgs)
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

func (h *Handler) GetFile(c *gin.Context) {
	logger := h.Logger.WithGroup("GetFile")

	file, ok := middleware.GetContextValue[model.MusmgrFile](c, middleware.CurrentFile)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx := c.Request.Context()

	rd, err := h.Storage.Read(ctx, file.ID)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rd.Close()

	c.Header("Content-Type", file.ContentType)
	c.Status(http.StatusOK)

	_, err = io.Copy(c.Writer, rd)
	if err != nil {
		logger.Info("Client disconnected")
		return
	}
}

func (h *Handler) createPreviewScore(ctx context.Context, id uuid.UUID, fileName string, pieceID uuid.UUID) (*uuid.UUID, error) {
	rd, err := h.Storage.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rd.Close()

	preview, err := h.PdfGenerator.Generate(ctx, rd)
	if err != nil {
		return nil, err
	}
	defer preview.Close()

	newID := uuid.New()

	err = h.Storage.Create(ctx, newID, preview, platform.UnknownSize, "application/pdf")
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

	if err = h.Queries.CreateFile(ctx, queryParams); err != nil {
		return nil, err
	}

	return &newID, nil
}

type createFileRequest struct {
	Classification string `form:"classification" binding:"required"`
	FileType       string `form:"file_type" binding:"required"`
	Name           string `form:"name" binding:"required"`
}

func (h *Handler) CreateFile(c *gin.Context) {
	logger := h.Logger.WithGroup("CreateFile")

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

	err = h.Queries.CreateFile(ctx, queryParams)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.Storage.Create(ctx, newFileID, file, fileHeader.Size, contentType)
	if err != nil {
		logger.Error(err.Error())

		delErr := h.Queries.DeleteFile(ctx, newFileID)
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
		previewID, err := h.createPreviewScore(ctx, newFileID, fileParams.Name, pieceID)
		if err != nil {
			logger.Error("didn't create preview file", "id", newFileID, "err", err)
			return
		}
		logger.Info("created preview", "id", newFileID, "preview_id", *previewID)
	}
}

type updateFileMetadataRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func (h *Handler) UpdateFileMetadata(c *gin.Context) {
	logger := h.Logger.WithGroup("UpdateFileMetadata")

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
		Description: textOrNull(req.Description),
		Name:        textOrNull(req.Name),
	}

	if err = h.Queries.UpdateFileMetadata(ctx, queryArgs); err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
	logger.Info("updated")
}

func (h *Handler) DeleteFile(c *gin.Context) {
	logger := h.Logger.WithGroup("DeleteFile")

	file, ok := middleware.GetContextValue[model.MusmgrFile](c, middleware.CurrentFile)
	if !ok {
		logger.Error("While converting file type")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx := c.Request.Context()

	err := h.Queries.DeleteFile(ctx, file.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Info("File doesn't exist")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.Storage.Delete(ctx, file.ID)
	if err != nil {
		logger.Warn("Failed to delete file", "id", file.ID, "err", err)
	}

	c.Status(http.StatusNoContent)
	logger.Info("Success")
}
