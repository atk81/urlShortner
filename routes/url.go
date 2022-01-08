package routes

import "github.com/gin-gonic/gin"
import "github.com/atk81/urlShortner/controllers"

func URLRoutes(router *gin.Engine) {
	url := router.Group("v1/short")
	{
		// Shorten a URL
		// curl -X POST -H "Content-Type: application/json" -d '{"url":"https://www.google.com"}' http://localhost:8080/v1/short
		url.POST("/", controllers.ShortenUrl)
		// Get the shortened URL
		// curl http://localhost:8080/v1/short/:shortUrl
		url.GET("/:shortUrl", controllers.GetUrl)
	}
}
