package websockets

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader upgrades HTTP requests to WebSocket connections.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// TODO: Replace this in production with your frontend origin(s).
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler handles websocket requests.
type Handler struct {
	Hub *Hub
}

// NewHandler creates a new websocket handler.
func NewHandler(hub *Hub) *Handler {
	return &Handler{
		Hub: hub,
	}
}

// ServeWS upgrades an authenticated HTTP request to a websocket connection.
func (h *Handler) ServeWS(c *gin.Context) {
	// UserID was set by JWTAuthMiddleware().
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("websocket upgrade failed: %v", err)
		return
	}

	client := &Client{
		UserID: userID,
		Conn:   conn,
		Hub:    h.Hub,
		Send:   make(chan Message, 256),
	}

	h.Hub.Register <- client

	log.Printf("User %d connected via WebSocket", userID)

	go client.WritePump()
	go client.ReadPump()
}
