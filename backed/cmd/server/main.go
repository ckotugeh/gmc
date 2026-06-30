package main

import (
	"log"

	"doctor-platform/internal/auth"
	"doctor-platform/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.ConnectDB()

	database.DB.AutoMigrate(&auth.User{})

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/auth/register", auth.Register)
		api.POST("/auth/login", auth.Login)
	}
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal(err)
	}
	router.Run(":8080")
}
