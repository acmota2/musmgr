package controller

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func (cc *ControllerContext) GetCategories(context *gin.Context) {
	categories, err := cc.Queries.GetCategories(cc.Context)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": categories})
	}
}

func (cc *ControllerContext) GetCategorySubcategories(context *gin.Context) {
	id := context.Param("name")
	matched, err := regexp.MatchString(`^[\w_\-]+$`, id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if !matched {
		context.JSON(http.StatusBadRequest, gin.H{"error": `the request was bad, you should feel bad`})
		return
	}

	subCats, err := cc.Queries.GetCategorySubcategories(cc.Context, id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": subCats})
	}
}
