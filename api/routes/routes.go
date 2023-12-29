package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	InitFileRoutes(router)
	InitUserRoutes(router)
	InitProjectRoutes(router)
}
