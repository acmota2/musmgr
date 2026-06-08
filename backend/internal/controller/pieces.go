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

func (h *BaseHandler) CreatePiece(c *gin.Context) {
	logger := h.Logger.WithGroup("CreatePiece")

	var req createPieceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := formatPieceDate(req.ComposedAt)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	newID := uuid.New()
	queryArgs := model.CreatePieceParams{
		ID:              newID,
		ComposedAt:      req.ComposedAt,
		Description:     req.Description,
		Instrumentation: req.Instrumentation,
		Title:           req.Title,
	}

	if err := h.Queries.CreatePiece(ctx, queryArgs); err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Location", fmt.Sprintf("/pieces/%s", newID.String()))
	c.Status(http.StatusCreated)
	logger.Info("success")
}

type updatePieceRequest struct {
	ComposedAt      *string                          `json:"composed_at"`
	Description     *string                          `json:"description"`
	Instrumentation *model.MusmgrInstrumentationName `json:"instrumentation"`
	Title           *string                          `json:"title"`
}

func (h *BaseHandler) UpdatePiece(c *gin.Context) {
	logger := h.Logger.WithGroup("UpdatePiece")

	pieceID, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var req updatePieceRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	_, err = formatPieceDate(*req.ComposedAt)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	queryArgs := model.UpdatePieceParams{
		ID:          pieceID,
		Description: TextOrNull(req.Description),
		Title:       TextOrNull(req.Title),
		ComposedAt:  TextOrNull(req.ComposedAt),
	}

	if req.Instrumentation != nil {
		queryArgs.Instrumentation = model.NullMusmgrInstrumentationName{
			MusmgrInstrumentationName: *req.Instrumentation,
			Valid:                     true,
		}
	}

	err = h.Queries.UpdatePiece(ctx, queryArgs)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
	logger.Info("success")
}

func (h *BaseHandler) DeletePiece(c *gin.Context) {
	group := "DeletePiece"
	logger := h.Logger.WithGroup(group)

	pieceID, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
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
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fileIDs := make([]uuid.UUID, 0, len(files))
	for _, file := range files {
		fileIDs = append(fileIDs, file.ID)
	}
	h.BestEffortDelete(group, fileIDs...)

	c.Status(http.StatusNoContent)
	logger.Info("success")
}

func (h *BaseHandler) GetPieces(c *gin.Context) {
	logger := h.Logger.WithGroup("GetPieces")

	ctx := c.Request.Context()

	pieces, err := h.Queries.GetPieces(ctx)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, pieces)
	logger.Info("success")
}

func (h *BaseHandler) GetPiece(c *gin.Context) {
	logger := h.Logger.WithGroup("GetPiece")

	pieceID, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	piece, err := h.Queries.GetPiece(ctx, pieceID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, piece)
	logger.Info("success")
}

func (h *BaseHandler) GetPieceEvents(c *gin.Context) {
	logger := h.Logger.WithGroup("GetPieceEvents")

	id, err := uuid.Parse(c.Param("piece_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	events, err := h.Queries.GetPieceEvents(ctx, id)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, events)
	logger.Info("success")
}
