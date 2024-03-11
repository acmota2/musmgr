package controller

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSongToEvent(context *gin.Context) {
	var songEvent model.SongEvent
	if err := context.ShouldBindJSON(&songEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	se, err := songEvent.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": se})
	}
}
