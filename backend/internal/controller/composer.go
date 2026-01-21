package controller

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) GetComposer(c *gin.Context) {
	ctx := c.Request.Context()
	composer, err := h.Queries.GetComposer(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, composer)
}

func (h *Handler) CreateComposer(c *gin.Context) {
	var req model.CreateComposerParams
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err := h.Queries.CreateComposer(ctx, req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Location", "/composer")
	c.Status(http.StatusCreated)
}

func (h *Handler) GetComposerPicture(c *gin.Context) {
	ctx := c.Request.Context()

	composer, err := h.Queries.GetComposer(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !composer.Picture.Valid {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	rd, err := h.Storage.Read(ctx, composer.Picture.Bytes)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	_, err = io.Copy(c.Writer, rd)
	if err != nil {
		log.Println("[INFO] Client disconnected, ignoring")
	}

	c.Status(http.StatusOK)
}

type updateComposerRequest struct {
	FullName  *string `json:"full_name"`
	Biography *string `json:"biography"`
}

func (h *Handler) UpdateComposer(c *gin.Context) {
	var req updateComposerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	queryArgs := model.UpdateComposerParams{
		FullName:  textOrNull(req.FullName),
		Biography: textOrNull(req.Biography),
	}
	if err := h.Queries.UpdateComposer(ctx, queryArgs); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Location", "/composer")
	c.Status(http.StatusNoContent)
}

func (h *Handler) UpdateComposerPicture(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "'file' is required"},
		)
		return
	}

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
		c.AbortWithStatus(http.StatusInternalServerError)
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

	err = h.DBTransaction(ctx, func(qtx *model.Queries) error {
		composer, err := qtx.GetComposer(ctx)
		if err != nil {
			return err
		}

		newPictureID := uuid.New()
		if err = h.Storage.Create(ctx, newPictureID, file, fileHeader.Size, contentType); err != nil {
			return err
		}

		queryArgs := model.UpdateComposerPictureParams{
			Picture:            pgtype.UUID{Bytes: newPictureID, Valid: true},
			PictureContentType: pgtype.Text{String: contentType, Valid: true},
		}
		if err = qtx.UpdateComposerPicture(ctx, queryArgs); err != nil {
			// try to delete
			h.BestEffortDelete(newPictureID)
			return err
		}
		if composer.Picture.Valid {
			h.BestEffortDelete(composer.Picture.Bytes)
		}

		return nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Location", "/composer/picture")
	c.Status(http.StatusNoContent)
}

func (h *Handler) DeleteComposerPicture(c *gin.Context) {
	ctx := c.Request.Context()
	pictureID, err := h.Queries.DeleteComposerPicture(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if pictureID.Valid {
		h.BestEffortDelete(pictureID.Bytes)
	}

	c.Status(http.StatusNoContent)
}
