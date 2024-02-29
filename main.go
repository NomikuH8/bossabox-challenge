package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.DebugMode)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env:", err)
	}

	g := gin.Default()

	port := ":" + os.Getenv("SERVER_PORT")

	g.Run(port)
}
