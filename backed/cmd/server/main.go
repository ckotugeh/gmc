package main

import (
	"log"

	"doctor-platform/internal/auth"
	"doctor-platform/internal/comments"
	"doctor-platform/internal/communities"
	"doctor-platform/internal/database"
	"doctor-platform/internal/middleware"
	"doctor-platform/internal/posts"
	"doctor-platform/internal/profile"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.ConnectDB()

	database.DB.AutoMigrate(
		&auth.User{},
		&profile.Profile{},
		&communities.Community{},
		&posts.Post{},
		&comments.Comment{},
	)

	router := gin.Default()

	// Communities module
	repo := communities.NewRepository()
	service := communities.NewService(repo)
	handler := communities.NewHandler(service)

	api := router.Group("/api")

	// Public routes
	api.POST("/auth/register", auth.Register)
	api.POST("/auth/login", auth.Login)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.GET("/me", func(c *gin.Context) {
		userID := c.GetUint("userID")
		email := c.GetString("email")

		c.JSON(200, gin.H{
			"user_id": userID,
			"email":   email,
		})
	})

	profile.RegisterRoutes(protected)
	communities.RegisterRoutes(protected, handler)
	posts.RegisterRoutes(router)
	comments.RegisterRoutes(router)

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal(err)
	}

	router.Run(":8080")
}
