package websockets

import (
	"github.com/gin-gonic/gin"

	"doctor-platform/internal/middleware"
)

// RegisterRoutes registers websocket routes.
func RegisterRoutes(router *gin.RouterGroup, hub *Hub) {
	handler := NewHandler(hub)

	ws := router.Group("/ws")
	ws.Use(middleware.JWTAuthMiddleware())
	{
		ws.GET("", handler.ServeWS)
	}
}
