package controller

import (
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSongSubCategories(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songs, err := model.GetSongSubcategories(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": songs})
	}
}

func GetAllSongs(context *gin.Context) {
	songs, err := model.GetAllSongs()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": songs})
	}
}

func CreateSong(context *gin.Context) {
	var song model.Song
	if err := context.ShouldBindJSON(&song); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saved, err := song.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": saved})
	}
}
