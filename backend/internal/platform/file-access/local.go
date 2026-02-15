package platform

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type localStorage struct {
	path string
}

func earlyCancel(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	return nil
}

func newLocalStorage(path string) (*localStorage, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}
	return &localStorage{path}, nil
}

func (ls *localStorage) makePath(name string) string {
	return filepath.Join(ls.path, name)
}

func (ls *localStorage) Create(ctx context.Context, id uuid.UUID, r io.Reader, size int64, contentType string) error {
	if err := earlyCancel(ctx); err != nil {
		return err
	}

	file, err := os.OpenFile(
		ls.makePath(id.String()),
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	if size >= 0 {
		_, err = io.CopyN(file, r, size)
	} else {
		_, err = io.Copy(file, r)
	}

	return err
}

func (ls *localStorage) Read(ctx context.Context, id uuid.UUID) (io.ReadCloser, error) {
	if err := earlyCancel(ctx); err != nil {
		return nil, err
	}

	return os.Open(ls.makePath(id.String()))
}

func (ls *localStorage) Delete(ctx context.Context, id uuid.UUID) error {
	if err := earlyCancel(ctx); err != nil {
		return err
	}

	return os.Remove(ls.makePath(id.String()))
}
