package controller

import (
	"backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cc *ControllerContext) CreateWorkEvent(context *gin.Context) {
	var workEvent model.CreateWorkEventParams
	if err := context.ShouldBindJSON(&workEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := cc.Queries.CreateWorkEvent(cc.Context, workEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": workEvent})
	}
}
