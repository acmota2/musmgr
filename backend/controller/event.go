package controller

import (
	"net/http"
	"strconv"

	"backend/model"

	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	var event model.Event
	if err := context.BindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": event})
	}
}

func GetAllSongsFromEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songs, err := model.GetEventSongs(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": songs})
	}
}

func GetAllEvents(context *gin.Context) {
	events, err := model.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": events})
	}
}
