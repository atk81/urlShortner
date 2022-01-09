package routes

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine{
	// Setup default router
	router := gin.Default()
	// Setup url routes
	URLRoutes(router)
	return router
}