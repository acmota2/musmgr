package controller

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SongCategories(context *gin.Context) {

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
