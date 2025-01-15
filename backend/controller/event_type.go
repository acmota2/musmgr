package controller

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (cc *ControllerContext) GetEventTypes(context *gin.Context) {
	eventTypes, err := cc.Queries.GetEventTypes(cc.Context)
	if errors.Is(err, pgx.ErrNoRows) {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": eventTypes})
	}
}

func (cc *ControllerContext) GetEventTypeEvents(context *gin.Context) {
	id := context.Param("name")
	matched, err := regexp.MatchString(`^[\w_\-][\w_\-\s]*$`, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if !matched {
		context.JSON(http.StatusBadRequest, gin.H{"error": `the request was bad, you should feel bad`})
		return
	}
	events, err := cc.Queries.GetEventTypeEvents(cc.Context, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": events})
	}
}
