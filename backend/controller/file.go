package controller

import (
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DownloadTextFile(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songFile, err := model.GetTextFileFromSong(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.FileAttachment(songFile.Path, songFile.Name+".json")
	}
}

func CreateSongFile(context *gin.Context) {
	var file model.SongFile
	if err := context.BindJSON(&file); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	saved, err := file.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": saved})
	}
}

func GetSongText(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	sf, err := model.GetTextFileFromSong(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.FileAttachment(sf.Path, sf.Name+".json")
	}
}

func GetAllFilesInformationFromSong(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, err := model.GetAllFilesFromSong(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": files})
	}
}
