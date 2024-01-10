package controller

import (
	"net/http"
	"regexp"

	"backend/model"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(context *gin.Context) {
	categories, err := model.GetAllCategories()

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": categories})
	}
}

func GetAllSubCategoriesFromCategory(context *gin.Context) {
	id := context.Param("name")
	matched, err := regexp.MatchString(`^[\w_\-]+$`, id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if !matched {
		context.JSON(http.StatusBadRequest, gin.H{"error": `the request was bad, you should feel bad`})
		return
	}

	subCats, err := model.GetAllSubCategoriesFromCategory(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": subCats})
	}
}
