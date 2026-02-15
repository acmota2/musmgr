package services

import (
	"context"
	"io"
)

type PdfGenerator interface {
	Generate(ctx context.Context, src io.Reader) (io.ReadCloser, error)
}
