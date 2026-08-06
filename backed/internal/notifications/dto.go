package notifications

type CreateNotificationRequest struct {
	UserID        uint   `json:"user_id" binding:"required"`
	ActorID       *uint  `json:"actor_id,omitempty"`
	Type          string `json:"type" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Message       string `json:"message" binding:"required"`
	ReferenceID   *uint  `json:"reference_id,omitempty"`
	ReferenceType string `json:"reference_type,omitempty"`
	ActionURL     string `json:"action_url,omitempty"`
}

type UpdateNotificationRequest struct {
	IsRead bool `json:"is_read"`
}

type NotificationResponse struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ActorID       *uint  `json:"actor_id,omitempty"`
	Type          string `json:"type"`
	Title         string `json:"title"`
	Message       string `json:"message"`
	ReferenceID   *uint  `json:"reference_id,omitempty"`
	ReferenceType string `json:"reference_type,omitempty"`
	IsRead        bool   `json:"is_read"`
	ActionURL     string `json:"action_url,omitempty"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
