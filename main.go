package main

import (
	"github.com/atk81/urlShortner/model"
	"github.com/atk81/urlShortner/routes"
)

func main() {
	// Setup default router
	router := routes.SetupRoutes()
	// setup database
	model.InitDatabase()
	// start the server
	router.Run(":8080")
}