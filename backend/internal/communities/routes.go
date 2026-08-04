package communities

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.RouterGroup, handler *Handler) {
	communities := route.Group("/communities")

	{
		communities.POST("", handler.CreateCommunity)
		communities.GET("", handler.GetCommunities)
		communities.GET("/:id", handler.GetCommunity)
		communities.PUT("/:id", handler.UpdateCommunity)
		communities.DELETE("/:id", handler.DeleteCommunity)
	}
}
