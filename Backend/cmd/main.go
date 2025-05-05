package main

import (
	"ecommerce/internal/db"
	"ecommerce/internal/handlers"
	"ecommerce/internal/kafka"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	db.InitDB()
	r := gin.Default()
	handlers.SetupRoutes(r)
	port := os.Getenv("PORT")
	r.Run(":" + port)
	go kafka.ConsumeOrderEvents()
}
