package controller

import (
	"github.com/gin-gonic/gin"

	database "backend/db"
	"backend/model"
	"net/http"
)

func CreateSubCategory(context *gin.Context) {
	var subcategory model.SubCategory

	if err := context.ShouldBindJSON(subcategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saved, err := subcategory.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"savedData": saved})
	}
}

func GetAllSongsFromSubcategory(context *gin.Context) {
	var subcategory model.SubCategory

	if err := context.BindJSON(subcategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var songs []*model.Song
	// se isto estiver certo, é um milagre
	err := database.PsqlDB.Select(&songs).
		InnerJoins("subcategory").
		Where("song_subcategory.subcategory_id = ?", subcategory.ID).Error

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"found": songs})
	}
}
