package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/repositories"
	"github.com/urosh1g/collab-editor/routes"
)

func main() {

	config := GetConfig()
	db := GetDatabase(&config)

	db.AutoMigrate(&models.File{})
	fileRepository := repositories.NewFileRepository(db)
	_ = fileRepository

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "running"})
	})
	router.GET("/files", func(c *gin.Context) {
		routes.GetFiles(c, fileRepository)
	})
	router.POST("/files", func(c *gin.Context) {
		routes.CreateFile(c, fileRepository)
	})
	router.DELETE("/files/:id", func(c *gin.Context) {
		routes.DeleteFile(c, fileRepository)
	})

	router.Run()
}
