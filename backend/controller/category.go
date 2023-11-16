package controller

import (
	"net/http"
	"strconv"

	"backend/model"

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

func GetAllSubCategoriesFromCategory(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	subCats, err := model.GetAllSubCategoriesFromCategory(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"subcategories": subCats})
	}
}
