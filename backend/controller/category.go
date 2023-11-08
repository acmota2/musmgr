package controller

import (
	"github.com/gin-gonic/gin"

	"backend/model"
	"net/http"
)

func CreateCategory(context *gin.Context) {
	var category model.Category

	if err := context.ShouldBindJSON(&category); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
