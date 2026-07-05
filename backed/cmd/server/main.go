package main

import (
	"log"

	"doctor-platform/internal/admin"
	"doctor-platform/internal/appointments"
	"doctor-platform/internal/auth"
	"doctor-platform/internal/comments"
	"doctor-platform/internal/communities"
	"doctor-platform/internal/database"
	"doctor-platform/internal/hospitals"
	"doctor-platform/internal/medicalrecords"
	"doctor-platform/internal/messages"
	"doctor-platform/internal/middleware"
	"doctor-platform/internal/notifications"
	"doctor-platform/internal/posts"
	"doctor-platform/internal/presence"
	"doctor-platform/internal/profile"
	"doctor-platform/internal/reactions"
	"doctor-platform/internal/search"
	"doctor-platform/internal/uploads"
	"doctor-platform/internal/video_consultations"
	"doctor-platform/internal/websockets"

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
		&reactions.Reaction{},
		&messages.Message{},
		&notifications.Notification{},
		&appointments.Appointment{},
		&medicalrecords.MedicalRecord{},
		&hospitals.Hospital{},
		&uploads.Upload{},
		&video_consultations.VideoConsultation{},
		&presence.Presence{},
		&search.Search{},
		&admin.Admin{},
	)

	router := gin.Default()

	// Communities module
	repo := communities.NewRepository()
	service := communities.NewService(repo)
	handler := communities.NewHandler(service)

	//reactions
	reactionRepo := reactions.NewRepository()
	reactionService := reactions.NewService(reactionRepo)
	reactionHandler := reactions.NewHandler(reactionService)

	//messaes
	messageRepo := messages.NewRepository()
	messageService := messages.NewService(messageRepo)
	messageHandler := messages.NewHandler(messageService)

	//notification
	notificationRepo := notifications.NewRepository()
	notificationService := notifications.NewService(notificationRepo)
	notificationHandler := notifications.NewHandler(notificationService)

	hub := websockets.NewHub()
	go hub.Run()

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
	reactions.RegisterRoutes(api, reactionHandler)
	messages.RegisterRoutes(api, messageHandler)
	notifications.RegisterRoutes(api, notificationHandler)
	websockets.RegisterRoutes(api, hub)
	presence.RegisterRoutes(api, database.DB)
	appointments.RegisterRoutes(api, database.DB)
	medicalrecords.RegisterRoutes(api, database.DB)
	hospitals.RegisterRoutes(api, database.DB)
	uploads.RegisterRoutes(api, database.DB)
	video_consultations.RegisterRoutes(api, database.DB)
	search.RegisterRoutes(api, database.DB)
	admin.RegisterRoutes(api, database.DB)

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal(err)
	}

	router.Run(":8080")
}
