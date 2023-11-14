package controller

import (
	database "backend/db"
	"backend/model"
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
)

func SongCategories(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func CreateSong(context *gin.Context) {
	var song model.Song

	if err := context.ShouldBindJSON(song); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saved, err := song.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"song": saved})
	}
}

func GetSongWithID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var song model.Song
	err = database.PsqlDB.First(&song, id).Error

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"song": song})
	}
}
