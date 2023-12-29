package main

import (
	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/database"
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/routes"
	"github.com/urosh1g/collab-editor/services"
)

func main() {
	database.Db.AutoMigrate(&models.File{}, &models.User{}, &models.Project{})
	router := gin.Default()

	{
		router.GET("/files", func(c *gin.Context) {
			routes.GetFiles(c, services.FilesService)
		})
		router.POST("/files", func(c *gin.Context) {
			routes.CreateFile(c, services.FilesService)
		})
		router.DELETE("/files/:id", func(c *gin.Context) {
			routes.DeleteFile(c, services.FilesService)
		})
	}

	{
		router.GET("/users", func(c *gin.Context) {
			routes.GetUsers(c, services.UsersService)
		})

		router.POST("/users", func(c *gin.Context) {
			routes.CreateUser(c, services.UsersService)
		})
	}

	{
		router.GET("/projects", func(c *gin.Context) {
			routes.GetProjects(c, services.ProjectsService)
		})

		router.POST("/projects", func(c *gin.Context) {
			routes.CreateProject(c, services.ProjectsService)
		})

		router.DELETE("/projects/:id", func(c *gin.Context) {
			routes.DeleteProject(c, services.ProjectsService)
		})
	}

	router.Run()
}
