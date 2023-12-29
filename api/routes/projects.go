package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/services"
)

func InitProjectRoutes(router *gin.Engine) {
	projectGroup := router.Group("/projects")
	projectGroup.GET("", func(c *gin.Context) {
		GetProjects(c, services.ProjectsService)
	})

	projectGroup.POST("", func(c *gin.Context) {
		CreateProject(c, services.ProjectsService)
	})

	projectGroup.DELETE("/:id", func(c *gin.Context) {
		DeleteProject(c, services.ProjectsService)
	})
}

func GetProjects(c *gin.Context, service *services.ProjectService) {
	projects, err := service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func CreateProject(c *gin.Context, service *services.ProjectService) {
	var projectRequest models.CreateProjectRequest
	if err := c.Bind(&projectRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	}
	project, err := service.Create(projectRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	}
	c.JSON(http.StatusOK, project)
}

func DeleteProject(c *gin.Context, service *services.ProjectService) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project with given id not found"})
		return
	}
	project, err := service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	}
	c.JSON(http.StatusOK, project)
}
