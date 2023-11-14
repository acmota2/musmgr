package controller

import (
	"net/http"
	"strconv"

	"backend/model"
	"backend/songs"
	"github.com/gin-gonic/gin"
)

func CreateSongFile(context *gin.Context) {
	var file model.SongFile

	if err := context.BindJSON(&file); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	saved, err := file.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"savedData": saved})
	}
}

func GetSongText[T songs.Chord](context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func GetAllFilesFromSong(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	files, err := model.GetAllFilesFromSong(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"found": files})
	}
}
