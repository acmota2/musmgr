package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (cc *ControllerContext) GetTextFile(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songFile, err := cc.Queries.GetTextFile(cc.Context, id.String())
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.FileAttachment(songFile.FilePath, songFile.Name)
	}
}

/*
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
*/

func (cc *ControllerContext) GetSongFiles(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, err := cc.Queries.GetSongFiles(cc.Context, id.String())
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": files})
	}
}
