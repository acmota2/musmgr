package controller

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (cc *ControllerContext) GetSongSubcategories(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songs, err := cc.Queries.GetSongSubcategories(cc.Context, id.String())
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": songs})
	}
}

func (cc *ControllerContext) GetSongs(context *gin.Context) {
	songs, err := cc.Queries.GetSongs(cc.Context)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": songs})
	}
}

func (cc *ControllerContext) CreateSong(context *gin.Context) {
	var song model.CreateSongParams
	if err := context.ShouldBindJSON(&song); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.Queries.CreateSong(cc.Context, song)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": song})
	}
}
