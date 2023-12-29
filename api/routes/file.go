package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/services"
)

func GetFiles(c *gin.Context, service *services.FileService) {
	result, err := service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed fetching files"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateFile(c *gin.Context, service *services.FileService) {
	var file models.CreateFileRequest
	if c.ShouldBind(&file) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid POST data"})
		return
	}
	result, err := service.Create(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteFile(c *gin.Context, service *services.FileService) {
	var file models.File
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid `id` path parameter }"})
		return
	}
	file, err = service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, file)
}
