package controller

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cc *ControllerContext) GetWorks(context *gin.Context) {
	songs, err := cc.Queries.GetWorks(cc.Context)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": songs})
	}
}

func (cc *ControllerContext) CreateWork(context *gin.Context) {
	var song model.CreateWorkParams
	if err := context.ShouldBindJSON(&song); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.Queries.CreateWork(cc.Context, song)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": song})
	}
}
