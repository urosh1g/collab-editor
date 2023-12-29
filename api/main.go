package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/repositories"
	"github.com/urosh1g/collab-editor/routes"
	"github.com/urosh1g/collab-editor/services"
)

func main() {
	config := GetConfig()
	db := GetDatabase(&config)

	db.AutoMigrate(&models.File{}, &models.User{}, &models.Project{})
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	fileRepository := repositories.NewFileRepository(db)

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

	router.GET("/users", func(c *gin.Context) {
		routes.GetUsers(c, userService)
	})

	router.POST("/users", func(c *gin.Context) {
		routes.CreateUser(c, userService)
	})

	router.Run()
}
