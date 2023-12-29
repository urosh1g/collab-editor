package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/services"
)

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
