package websockets

import (
	"doctor-platform/internal/websockets"
	"net/http"

	"github.com/gin-gonic/gin"
)

var upgrader = websockets.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			break
		}

		conn.WriteMessage(websockets.TextMessage, msg)
	}
}
