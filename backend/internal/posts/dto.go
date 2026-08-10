package posts

// ----------------------------------------
// Create Post
// ----------------------------------------

type CreatePostRequest struct {
	CommunityID uint   `json:"community_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`

	// Optional
	ContentType string `json:"content_type"`
	ImageURL    string `json:"image_url"`

	IsAnonymous bool `json:"is_anonymous"`

	Tags        []string            `json:"tags"`
	Attachments []AttachmentRequest `json:"attachments"`
	Poll        *CreatePollRequest  `json:"poll"`
}

// ----------------------------------------
// Update Post
// ----------------------------------------

type UpdatePostRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`

	ImageURL string `json:"image_url"`

	IsAnonymous *bool `json:"is_anonymous"`
	IsPinned    *bool `json:"is_pinned"`
	IsLocked    *bool `json:"is_locked"`

	Tags        []string            `json:"tags"`
	Attachments []AttachmentRequest `json:"attachments"`
	Poll        *UpdatePollRequest  `json:"poll"`
}

// ----------------------------------------
// Attachment DTO
// ----------------------------------------

type AttachmentRequest struct {
	FileName string `json:"file_name"`
	FileURL  string `json:"file_url"`
	FileType string `json:"file_type"`
	FileSize int64  `json:"file_size"`
}

// ----------------------------------------
// Poll DTO
// ----------------------------------------

type CreatePollRequest struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

type UpdatePollRequest struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

// ----------------------------------------
// Response DTO
// ----------------------------------------

type PostResponse struct {
	ID uint `json:"id"`

	CommunityID uint `json:"community_id"`
	AuthorID    uint `json:"author_id"`

	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`

	ImageURL string `json:"image_url"`

	IsAnonymous bool `json:"is_anonymous"`
	IsEdited    bool `json:"is_edited"`
	IsPinned    bool `json:"is_pinned"`
	IsLocked    bool `json:"is_locked"`

	LikesCount    int `json:"likes_count"`
	CommentsCount int `json:"comments_count"`
	ViewsCount    int `json:"views_count"`

	Attachments []Attachment `json:"attachments"`
	Tags        []Tag        `json:"tags"`
	Polls       []Poll       `json:"polls"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
