package reactions

type CreateReactionRequest struct {
	ReactionType string `json:"reaction_type" binding:"required,oneof=like dislike"`
}

type UpdateReactionRequest struct {
	ReactionType string `json:"reaction_type" binding:"required,oneof=like dislike"`
}

type ReactionResponse struct {
	ID           uint   `json:"id"`
	PostID       uint   `json:"post_id"`
	UserID       uint   `json:"user_id"`
	ReactionType string `json:"reaction_type"`
}
