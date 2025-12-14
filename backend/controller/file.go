package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


/*
func CreateWorkFile(context *gin.Context) {
	var file model.WorkFile
	if err := context.BindJSON(&file); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	saved, err := file.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"data": saved})
	}
}

func GetWorkText(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	sf, err := model.GetTextFileFromWork(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.FileAttachment(sf.Path, sf.Name+".json")
	}
}
*/

func (cc *ControllerContext) GetWorkFiles(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, err := cc.Queries.GetWorkFiles(cc.Context, id.String())
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		// TODO: Decide between MinIO or serving directly from server
		context.JSON(http.StatusOK, gin.H{"data": files})
	}
}
