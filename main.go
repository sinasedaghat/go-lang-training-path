package main

import (
	"github.com/gin-gonic/gin"

	"sample-url.com/REST-API/database"
	"sample-url.com/REST-API/routes"
)

func main() {
	database.Initialize()
	server := gin.Default()
	// defer server.Run(":8080") // Don't call it without "defer" because none of the endpoints are registered. Or call it at the end of the function.
	// You should not defer server.Run() — this will delay your server start until the main() function exits.

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
