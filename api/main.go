package main

import (
	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/database"
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/routes"
)

func main() {
	database.Db.AutoMigrate(&models.File{}, &models.User{}, &models.Project{})
	router := gin.Default()
	routes.InitRoutes(router)
	router.Run()
}
