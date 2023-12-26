package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/repositories"
	"github.com/urosh1g/collab-editor/models"
	"net/http"
	"strconv"
)

func GetFiles(c *gin.Context, repo repositories.Repository[models.File]) {
	result, err := repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed fetching files"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateFile(c *gin.Context, repo repositories.Repository[models.File]) {
	var file models.File
	if c.ShouldBind(&file) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid POST data"})
		return
	}
	result, err := repo.Create(&file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteFile(c *gin.Context, repo repositories.Repository[models.File]) {
	var file models.File
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid `id` path parameter }"})
		return
	}
	file, err = repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, file)
}
