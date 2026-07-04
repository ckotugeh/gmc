package websockets

// Service defines websocket business logic layer.
type Service struct {
	hub *Hub
}

// NewService creates websocket service.
func NewService(hub *Hub) *Service {
	return &Service{
		hub: hub,
	}
}

// BroadcastMessage sends a message to a specific user.
func (s *Service) BroadcastMessage(userID uint, msg Message) {
	s.hub.Broadcast <- msg
}
