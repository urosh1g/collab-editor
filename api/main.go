package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Failed to connect to DB %v", err))
	}

	db.AutoMigrate(&models.File{})
	fileRepository := repositories.NewFileRepository(db)
	_ = fileRepository

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "running"})
	})

	router.Run()
}
