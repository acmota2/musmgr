package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/policies"
	"github.com/jackc/pgx/v5"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func formatPieceDate(s string) (time.Time, error) {
	return time.Parse("2006", s)
}

type createPieceRequest struct {
	ComposedAt      string                          `json:"composed_at"`
	Description     string                          `json:"description"`
	Instrumentation model.MusmgrInstrumentationName `json:"instrumentation"`
	Title           string                          `json:"title"`
}

func (h *Handler) CreatePiece(c *gin.Context) {
	var req createPieceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err := formatPieceDate(req.ComposedAt)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	newID := uuid.New()
	queryArgs := model.CreatePieceParams{
		ID:              newID,
		ComposedAt:      req.ComposedAt,
		Instrumentation: req.Instrumentation,
		Title:           req.Title,
	}

	if err := h.Queries.CreatePiece(ctx, queryArgs); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Location", fmt.Sprintf("/pieces/%s", newID.String()))
	c.Status(http.StatusCreated)
}

type updatePieceRequest struct {
	ComposedAt      *string                          `json:"composed_at"`
	Description     *string                          `json:"description"`
	Instrumentation *model.MusmgrInstrumentationName `json:"instrumentation"`
	Title           *string                          `json:"title"`
}

func (h *Handler) UpdatePiece(c *gin.Context) {
	pieceID, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req updatePieceRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err = formatPieceDate(*req.ComposedAt)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	queryArgs := model.UpdatePieceParams{
		ID:          pieceID,
		Description: textOrNull(req.Description),
		Title:       textOrNull(req.Title),
		ComposedAt:  textOrNull(req.ComposedAt),
	}

	if req.Instrumentation != nil {
		queryArgs.Instrumentation = model.NullMusmgrInstrumentationName{
			MusmgrInstrumentationName: *req.Instrumentation,
			Valid:                     true,
		}
	}

	err = h.Queries.UpdatePiece(ctx, queryArgs)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) DeletePiece(c *gin.Context) {
	group := "DeletePiece"

	pieceID, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	var files []model.MusmgrFile
	err = h.DBTransaction(ctx, func(qtx *model.Queries) error {
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
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	fileIDs := make([]uuid.UUID, 0, len(files))
	for _, file := range files {
		fileIDs = append(fileIDs, file.ID)
	}
	h.BestEffortDelete(group, fileIDs...)

	c.Status(http.StatusNoContent)
}

func (h *Handler) GetPieces(c *gin.Context) {
	ctx := c.Request.Context()

	pieces, err := h.Queries.GetPieces(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, pieces)
}

func (h *Handler) GetPiece(c *gin.Context) {
	pieceID, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	piece, err := h.Queries.GetPiece(ctx, pieceID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, piece)
}

func (h *Handler) GetPieceEvents(c *gin.Context) {
	id, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	events, err := h.Queries.GetPieceEvents(ctx, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, events)
}
