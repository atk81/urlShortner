package routes

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine{
	// Setup default router
	router := gin.Default()
	return router
}