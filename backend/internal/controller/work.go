package controller

import (
	"backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cc *ControllerContext) GetWorks(context *gin.Context) {
	works, err := cc.Queries.GetWorks(cc.Context)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": works})
	}
}

func (cc *ControllerContext) CreateWork(context *gin.Context) {
	var work model.CreateWorkParams
	if err := context.ShouldBindJSON(&work); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.Queries.CreateWork(cc.Context, work)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": work})
	}
}
