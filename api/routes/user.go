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

func GetUsers(c *gin.Context, repo repositories.Repository[models.User]) {
	result, err := repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed fetching Users"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context, service *services.UserService) {
	var User models.User
	if c.ShouldBind(&User) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid POST data"})
		return
	}
	fmt.Printf("%+v", User)
	result, err := service.Create(&User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, result)
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