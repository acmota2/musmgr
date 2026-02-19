package controller

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/middleware"
	"github.com/acmota2/musmgr/backend/internal/model"
	platform "github.com/acmota2/musmgr/backend/internal/platform/file-access"
	"github.com/acmota2/musmgr/backend/internal/policies"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) GetPieceFiles(c *gin.Context) {
	pieceId, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, files)
}

func (h *Handler) GetFile(c *gin.Context) {
	file, ok := middleware.GetContextValue[model.MusmgrFile](c, middleware.CurrentFile)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx := c.Request.Context()

	rd, err := h.Storage.Read(ctx, file.ID)
	if err != nil {
		log.Printf("Failed here with error %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rd.Close()

	c.Header("Content-Type", file.ContentType)
	c.Status(http.StatusOK)

	_, err = io.Copy(c.Writer, rd)
	if err != nil {
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
	Classification string               `form:"classification" binding:"required"`
	FileType       model.MusmgrFileType `form:"file_type" binding:"required"`
	Name           string               `form:"name" binding:"required"`
}

func (h *Handler) CreateFile(c *gin.Context) {
	pieceID, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var fileParams createFileRequest
	if err = c.ShouldBind(&fileParams); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// for early return
	class, err := policies.StringToClassification(fileParams.Classification)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newFileID := uuid.New()
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
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
		FileType:       fileParams.FileType,
		Name:           fileParams.Name,
		Origin:         model.MusmgrFileOriginUser,
		ParentID:       pgtype.UUID{Valid: false},
		PieceID:        pieceID,
	}

	err = h.Queries.CreateFile(ctx, queryParams)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.Storage.Create(ctx, newFileID, file, fileHeader.Size, contentType)
	if err != nil {
		_ = h.Queries.DeleteFile(ctx, newFileID)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Location", fmt.Sprintf("/pieces/%s/files/%s", pieceID, newFileID))
	c.JSON(http.StatusCreated, gin.H{
		"id": newFileID,
	})

	if fileParams.FileType == model.MusmgrFileTypeScoreFull {
		previewID, err := h.createPreviewScore(ctx, newFileID, fileParams.Name, pieceID)
		if err != nil {
			log.Printf("ERR: didn't create preview file for %s with error: %v", newFileID, err)
			return
		}
		log.Printf("INFO: Created preview with id: %s", *previewID)
	}
}

type updateFileMetadataRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func (h *Handler) UpdateFileMetadata(c *gin.Context) {
	fileID, err := uuid.Parse(c.Param("file_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	var req updateFileMetadataRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	queryArgs := model.UpdateFileMetadataParams{
		ID:          fileID,
		Description: textOrNull(req.Description),
		Name:        textOrNull(req.Name),
	}

	if err = h.Queries.UpdateFileMetadata(ctx, queryArgs); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) DeleteFile(c *gin.Context) {
	file, ok := middleware.GetContextValue[model.MusmgrFile](c, middleware.CurrentFile)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx := c.Request.Context()

	err := h.Queries.DeleteFile(ctx, file.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.Storage.Delete(ctx, file.ID)
	if err != nil {
		log.Printf("WARN: failed to delete file %s: %v", file.ID, err)
	}

	c.Status(http.StatusNoContent)
}
