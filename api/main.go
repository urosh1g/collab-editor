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
	fileRepository := repositories.NewFileRepository(db)
	userRepository := repositories.NewUserRepository(db)
	projectRepository := repositories.NewProjectRepository(db)
	userService := services.NewUserService(userRepository)
	projectService := services.NewProjectService(projectRepository)

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

	router.GET("/projects", func(c *gin.Context) {
		routes.GetProjects(c, projectService)
	})

	router.POST("/projects", func(c *gin.Context) {
		routes.CreateProject(c, projectService)
	})

	router.DELETE("/projects/:id", func(c *gin.Context) {
		routes.DeleteProject(c, projectService)
	})

	router.Run()
}
