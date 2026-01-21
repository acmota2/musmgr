package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/middleware"
	"github.com/acmota2/musmgr/backend/internal/model"
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
	file, ok := middleware.GetContextValue[model.MusmgrFile](c, middleware.ClassKey)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx := c.Request.Context()

	rd, err := h.Storage.Read(ctx, file.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rd.Close()

	c.Header("Content-Type", "application/octet-stream")
	c.Status(http.StatusOK)

	_, err = io.Copy(c.Writer, rd)
	if err != nil {
		return
	}
}

type createFileRequest struct {
	Classification policies.FileClassification `form:"classification" binding:"required"`
	FileType       model.MusmgrFileType        `form:"file_type" binding:"required"`
	Name           string                      `form:"name" binding:"required"`
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

	ctx := c.Request.Context()

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

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	queryParams := model.CreateFileParams{
		ID:             newFileID,
		Classification: int16(fileParams.Classification),
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
	file, ok := middleware.GetContextValue[model.MusmgrFile](c, middleware.ClassKey)
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
