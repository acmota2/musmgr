package controller

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (cc *ControllerContext) CreateEvent(context *gin.Context) {
	var event model.CreateEventParams
	if err := context.BindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.Queries.CreateEvent(cc.Context, event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": event})
	}
}

func (cc *ControllerContext) GetEventWorks(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	works, err := cc.Queries.GetEventWorks(cc.Context, id.String())

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": works})
	}
}
