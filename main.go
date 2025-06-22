package main

import (
	"dts-developer-challenge/config"
	"dts-developer-challenge/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.Use(cors.Default()) // Enables CORS for all headers and origins

	routes.RegisterTaskRoutes(r)

	r.Run(":8080")
}
