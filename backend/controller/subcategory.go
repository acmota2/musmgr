package controller

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (cc *ControllerContext) CreateSubcategory(context *gin.Context) {
	var subcategory model.CreateSubcategoryParams
	if err := context.ShouldBindJSON(&subcategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.Queries.CreateSubcategory(cc.Context, subcategory)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": subcategory})
	}
}

func (cc *ControllerContext) GetSubcategorySongs(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songs, err := cc.Queries.GetSubcategorySongs(cc.Context, id.String())
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": songs})
	}
}
