package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nomikuh8/bossabox-challenge/src/routes"
)

func main() {
	gin.SetMode(gin.DebugMode)

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Starting server with default values for .env")
	}

	g := gin.Default()

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "3000"
	}

	routes.RegisterToolRoutes(g)

	g.Run(":" + port)
}
