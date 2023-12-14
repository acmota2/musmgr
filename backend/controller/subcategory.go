package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/model"
	"net/http"
)

func CreateSubCategory(context *gin.Context) {
	var subcategory model.SubCategory
	if err := context.ShouldBindJSON(&subcategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subCat, err := subcategory.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": subCat})
	}
}

func GetAllSongsFromSubcategory(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songs, err := model.GetAllSongsFromSubcategory(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": songs})
	}
}

func GetAllSubcategories(context *gin.Context) {
	subCats, err := model.GetAllSubCategories()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": subCats})
	}
}
