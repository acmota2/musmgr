package controller

import (
	database "backend/db"
	"backend/model"
	"backend/songs"

	"github.com/gin-gonic/gin"

	"net/http"
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
	var song model.Song

	if err := context.BindJSON(&song); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var files []*model.SongFile
	err := database.PsqlDB.Where("song_id = ?", song.ID).Find(&files).Error

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"found": files})
	}
}
