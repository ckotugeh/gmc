package main

import (
	"log"
	"os"
	"strings"

	"doctor-platform/internal/admin"
	"doctor-platform/internal/allergies"
	"doctor-platform/internal/appointments"
	"doctor-platform/internal/auth"
	"doctor-platform/internal/availability"
	"doctor-platform/internal/comments"
	"doctor-platform/internal/communities"
	"doctor-platform/internal/database"
	"doctor-platform/internal/doctor_reviews"
	"doctor-platform/internal/doctor_schedules"
	"doctor-platform/internal/hospitals"
	"doctor-platform/internal/lab_requests"
	"doctor-platform/internal/lab_results"
	"doctor-platform/internal/medical_specialties"
	"doctor-platform/internal/medicalrecords"
	"doctor-platform/internal/messages"
	"doctor-platform/internal/middleware"
	"doctor-platform/internal/notifications"
	"doctor-platform/internal/payments"
	"doctor-platform/internal/posts"
	"doctor-platform/internal/prescriptions"
	"doctor-platform/internal/presence"
	"doctor-platform/internal/profile"
	"doctor-platform/internal/reactions"
	"doctor-platform/internal/search"
	"doctor-platform/internal/uploads"
	"doctor-platform/internal/video_consultations"
	"doctor-platform/internal/vitals"
	"doctor-platform/internal/websockets"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
    log.Printf("No .env file found; using environment variables")
}
	// Connect database
	database.ConnectDB()

	// Auto migrate all models
	if err := database.DB.AutoMigrate(
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
		&payments.Payment{},
		&doctor_schedules.DoctorSchedule{},
		&availability.Availability{},
		&medical_specialties.MedicalSpecialty{},
		&doctor_reviews.DoctorReview{},
		&prescriptions.Prescription{},
		&prescriptions.PrescriptionItem{},
		&vitals.Vital{},
		&allergies.Allergy{},
		&lab_requests.LabRequest{},
		&lab_results.LabResult{},
	); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	// Build allowed origins from env (comma-separated) + always allow localhost for dev
	allowedOrigins := []string{"http://localhost:5173", "http://localhost:3000"}
	if raw := os.Getenv("ALLOWED_ORIGINS"); raw != "" {
		for _, o := range strings.Split(raw, ",") {
			o = strings.TrimSpace(o)
			if o != "" {
				allowedOrigins = append(allowedOrigins, o)
			}
		}
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
	}))

	// Trust proxy
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal(err)
	}

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":     "Doctor Platform API",
			"	version": "1.0.0",
			"status":   "running",
		})
	})

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// =========================
	// Dependency Injection
	// =========================

	communityRepo := communities.NewRepository()
	communityService := communities.NewService(communityRepo)
	communityHandler := communities.NewHandler(communityService)

	reactionRepo := reactions.NewRepository()
	reactionService := reactions.NewService(reactionRepo)
	reactionHandler := reactions.NewHandler(reactionService)

	messageRepo := messages.NewRepository()
	messageService := messages.NewService(messageRepo)
	messageHandler := messages.NewHandler(messageService)

	notificationRepo := notifications.NewRepository()
	notificationService := notifications.NewService(notificationRepo)
	notificationHandler := notifications.NewHandler(notificationService)

	// WebSocket Hub
	hub := websockets.NewHub()
	go hub.Run()

	// API group
	api := router.Group("/api")

	// Public routes
	api.POST("/auth/register", auth.Register)
	api.POST("/auth/login", auth.Login)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user_id": c.GetUint("userID"),
			"email":   c.GetString("email"),
		})
	})

	// Protected modules
	profile.RegisterRoutes(protected)
	communities.RegisterRoutes(protected, communityHandler)

	// API modules
	posts.RegisterRoutes(protected)
	comments.RegisterRoutes(protected)
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
	payments.RegisterRoutes(api, database.DB)
	doctor_schedules.RegisterRoutes(api, database.DB)
	availability.RegisterRoutes(api, database.DB)
	medical_specialties.RegisterRoutes(api, database.DB)
	doctor_reviews.RegisterRoutes(api, database.DB)
	prescriptions.RegisterRoutes(api, database.DB)
	vitals.RegisterRoutes(api, database.DB)
	allergies.RegisterRoutes(api, database.DB)
	lab_requests.RegisterRoutes(api, database.DB)
	lab_results.RegisterRoutes(api, database.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Doctor Platform API running on :%s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
