package websockets

// Repository defines websocket persistence layer (optional for now).
type Repository interface {
	SaveMessage(msg Message) error
	GetMessages(userID uint) ([]Message, error)
}
