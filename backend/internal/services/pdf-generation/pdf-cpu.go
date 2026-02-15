package services

import (
	"bytes"
	"context"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type pdfcpuGenerator struct{}

func readerToReadSeeker(r io.Reader) (io.ReadSeeker, error) {
	if rs, ok := r.(io.ReadSeeker); ok {
		return rs, nil
	}

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}
	return bytes.NewReader(buf.Bytes()), nil
}

// might change if temp files are needed - in memory should be enough, however
func (gen pdfcpuGenerator) Generate(ctx context.Context, rd io.Reader) (io.ReadCloser, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	rds, err := readerToReadSeeker(rd)
	if err != nil {
		return nil, err
	}

	var wrTrim bytes.Buffer
	if err = api.Trim(rds, &wrTrim, []string{"1"}, model.NewDefaultConfiguration()); err != nil {
		return nil, err
	}

	wm, err := api.TextWatermark(
		"Preview",
		"font:Helvetica,align:c,pos:c,off:0 0,scale:1 rel,points:24,op:.8",
		true,
		false,
		types.POINTS,
	)
	if err != nil {
		return nil, err
	}

	wmIn := bytes.NewReader(wrTrim.Bytes())
	var wrWatermark bytes.Buffer
	if err = api.AddWatermarks(wmIn, &wrWatermark, nil, wm, model.NewDefaultConfiguration()); err != nil {
		return nil, err
	}

	rc := io.NopCloser(bytes.NewReader(wrWatermark.Bytes()))
	return rc, nil
}
