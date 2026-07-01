package main

import (
	"log"

	"doctor-platform/internal/auth"
	"doctor-platform/internal/database"
	"doctor-platform/internal/middleware"

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
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		protected.GET("/me", func(c *gin.Context) {
			userID := c.GetUint("userID")
			email := c.GetString("email")

			c.JSON(200, gin.H{
				"user_id": userID,
				"email":   email,
			})
		})
	}
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal(err)
	}
	router.Run(":8080")
}
