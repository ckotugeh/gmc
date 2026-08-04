package websockets

import (
	"log"
	"time"
)

// Hub manages all active websocket clients.
type Hub struct {
	// Connected clients keyed by user ID.
	Clients map[uint]*Client

	// Register requests from clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	// Messages to deliver.
	Broadcast chan Message
}

// NewHub creates a new websocket hub.
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[uint]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

// Run starts the websocket hub.
func (h *Hub) Run() {
	for {
		select {

		// Register a client.
		case client := <-h.Register:
			h.Clients[client.UserID] = client

			log.Printf("User %d connected", client.UserID)

			// Notify everyone that this user is online.
			h.broadcastPresence(client.UserID, true)

		// Unregister a client.
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, client.UserID)
				close(client.Send)

				log.Printf("User %d disconnected", client.UserID)

				// Notify everyone that this user went offline.
				h.broadcastPresence(client.UserID, false)
			}

		// Deliver a message.
		case msg := <-h.Broadcast:

			// Direct message.
			if msg.ReceiverID != 0 {
				if receiver, ok := h.Clients[msg.ReceiverID]; ok {
					select {
					case receiver.Send <- msg:
					default:
						close(receiver.Send)
						delete(h.Clients, receiver.UserID)
					}
				}
				continue
			}

			// Broadcast to everyone.
			for _, client := range h.Clients {
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(h.Clients, client.UserID)
				}
			}
		}
	}
}

// broadcastPresence informs connected users when someone
// comes online or goes offline.
func (h *Hub) broadcastPresence(userID uint, online bool) {
	msg := Message{
		Type:      MessageTypePresence,
		UserID:    userID,
		IsOnline:  online,
		Timestamp: time.Now(),
	}

	for _, client := range h.Clients {
		select {
		case client.Send <- msg:
		default:
			close(client.Send)
			delete(h.Clients, client.UserID)
		}
	}
}
