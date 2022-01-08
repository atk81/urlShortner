package controllers

import (
	"github.com/atk81/urlShortner/model"
	"github.com/atk81/urlShortner/utils"
	"github.com/gin-gonic/gin"

)

func ShortenUrl(c *gin.Context) {
	var urlShortner model.UrlShortner
	c.BindJSON(&urlShortner)
	longUrl := urlShortner.LongURL
	id, err := model.AddNewUrl(longUrl)
	if err!=nil{
		c.JSON(500,gin.H{
			"Message": err,
		})
		c.Abort()
	} 
	// If Added to database successfully.
	// Convert the interger to base62
	shortUrl := utils.Base62(id)
	c.JSON(200, gin.H{
		"shortUrl": shortUrl,
	})
}

func GetUrl(c *gin.Context){
	shortUrl := c.Param("shortUrl")
	// Convert a shortUrl(base62) to base10
	id := utils.Base10(shortUrl)
	longUrl, err := model.GetLongURL(id);
	if err!=nil{
		c.JSON(500,gin.H{
			"Message": err,
		})
		c.Abort()
	}
	c.JSON(200,gin.H{
		"longUrl": longUrl,
	})
}
