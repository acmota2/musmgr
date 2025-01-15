package controller

import (
	"backend/model"
	"context"
)

type ControllerContext struct {
	Queries *model.Queries
	Context context.Context
}

func New(queries *model.Queries, ctx context.Context) *ControllerContext {
	return &ControllerContext{queries, ctx}
}
