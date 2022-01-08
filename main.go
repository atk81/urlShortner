package main

import (
	"github.com/atk81/urlShortner/model"
	"github.com/gin-gonic/gin"
)

func main() {
	// setup default gin router
	router := gin.Default()
	// Create DB connection
	_ = model.InitDatabase()
	// setup routes
	setupRoutes(router)
	// start the server
	router.Run(":8080")
}


func setupRoutes(router *gin.Engine) {
	// setup routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
			"status":  "ok",
		})
	})
}