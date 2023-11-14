package controller

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(context *gin.Context) {
	categories, err := model.GetAllCategories()

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"found": categories})
	}
}
