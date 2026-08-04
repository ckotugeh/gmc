package search

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all search routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	search := rg.Group("/search")
	search.Use(middleware.JWTAuthMiddleware())
	{
		// Search history
		search.POST("", handler.CreateSearch)
		search.GET("", handler.GetMySearches)
		search.GET("/:id", handler.GetSearch)
		search.DELETE("/:id", handler.DeleteSearch)

		// Global search
		search.GET("/global", handler.GlobalSearch)

		// Category searches
		search.GET("/doctors", handler.SearchDoctors)
		search.GET("/patients", handler.SearchPatients)
		search.GET("/hospitals", handler.SearchHospitals)
		search.GET("/communities", handler.SearchCommunities)
		search.GET("/posts", handler.SearchPosts)
	}
}
