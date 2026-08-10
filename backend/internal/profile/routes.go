package profile

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {

	repository := NewRepository()
	service := NewService(repository)
	handler := NewHandler(service)

	router.POST("/profile", handler.CreateProfile)
	router.GET("/profile", handler.GetProfile)
	router.PUT("/profile", handler.UpdateProfile)
}
