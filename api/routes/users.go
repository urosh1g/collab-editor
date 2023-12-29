package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/repositories"
	"github.com/urosh1g/collab-editor/services"
)

func InitUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	userGroup.GET("", func(c *gin.Context) {
		GetUsers(c, services.UsersService)
	})
	userGroup.POST("", func(c *gin.Context) {
		CreateUser(c, services.UsersService)
	})
}

func GetUsers(c *gin.Context, service *services.UserService) {
	users, err := service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context, service *services.UserService) {
	var createRequest models.CreateUserRequest
	if err := c.Bind(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	}
	user, err := service.Create(createRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context, repo repositories.Repository[models.User]) {
	var User models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid `id` path parameter }"})
		return
	}
	User, err = repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, User)
}
