package services

import (
	"context"
	"io"
)

type PdfGenerator interface {
	Generate(ctx context.Context, src io.Reader) (io.ReadCloser, error)
}

type PdfGeneratorType int

const (
	PDFCPU = iota
)

// future proofing
func NewPdfGenerator(_gt PdfGeneratorType) PdfGenerator {
	return &pdfcpuGenerator{}
}
