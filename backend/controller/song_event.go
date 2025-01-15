package controller

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cc *ControllerContext) CreateSongEvent(context *gin.Context) {
	var songEvent model.CreateSongEventParams
	if err := context.ShouldBindJSON(&songEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := cc.Queries.CreateSongEvent(cc.Context, songEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": songEvent})
	}
}
