package messages

type CreateMessageRequest struct {
	ReceiverID    uint   `json:"receiver_id" binding:"required"`
	Content       string `json:"content" binding:"required"`
	MessageType   string `json:"message_type,omitempty"`
	AttachmentURL string `json:"attachment_url,omitempty"`
	ReplyToID     *uint  `json:"reply_to_id,omitempty"`
}

type UpdateMessageRequest struct {
	Content       string `json:"content" binding:"required"`
	AttachmentURL string `json:"attachment_url,omitempty"`
}

type MarkAsReadRequest struct {
	IsRead bool `json:"is_read"`
}

type MessageResponse struct {
	ID            uint   `json:"id"`
	SenderID      uint   `json:"sender_id"`
	ReceiverID    uint   `json:"receiver_id"`
	Content       string `json:"content"`
	MessageType   string `json:"message_type"`
	AttachmentURL string `json:"attachment_url"`
	ReplyToID     *uint  `json:"reply_to_id,omitempty"`
	IsRead        bool   `json:"is_read"`
	IsEdited      bool   `json:"is_edited"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
