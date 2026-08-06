package posts

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Service interface {
	CreatePost(req *CreatePostRequest, authorID uint) (*Post, error)
	GetPosts() ([]Post, error)
	GetPost(id uint) (*Post, error)
	GetCommunityPosts(communityID uint) ([]Post, error)
	UpdatePost(id uint, authorID uint, req *UpdatePostRequest) (*Post, error)
	DeletePost(id uint, authorID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// ----------------------------------------------------
// Create Post
// ----------------------------------------------------

func (s *service) CreatePost(req *CreatePostRequest, authorID uint) (*Post, error) {

	title := strings.TrimSpace(req.Title)
	content := strings.TrimSpace(req.Content)

	if title == "" {
		return nil, errors.New("title is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if req.CommunityID == 0 {
		return nil, errors.New("community_id is required")
	}

	contentType := strings.TrimSpace(req.ContentType)
	if contentType == "" {
		contentType = "markdown"
	}

	post := &Post{
		CommunityID: req.CommunityID,
		AuthorID:    authorID,
		Title:       title,
		Content:     content,
		ContentType: contentType,
		ImageURL:    req.ImageURL,
		IsAnonymous: req.IsAnonymous,
	}

	if err := s.repo.Create(post); err != nil {
		return nil, err
	}

	return post, nil
}

// ----------------------------------------------------
// Get All Posts
// ----------------------------------------------------

func (s *service) GetPosts() ([]Post, error) {
	return s.repo.GetAll()
}

// ----------------------------------------------------
// Get Single Post
// ----------------------------------------------------

func (s *service) GetPost(id uint) (*Post, error) {
	return s.repo.GetByID(id)
}

// ----------------------------------------------------
// Get Community Posts
// ----------------------------------------------------

func (s *service) GetCommunityPosts(communityID uint) ([]Post, error) {
	return s.repo.GetByCommunityID(communityID)
}

// ----------------------------------------------------
// Update Post
// ----------------------------------------------------

func (s *service) UpdatePost(id uint, authorID uint, req *UpdatePostRequest) (*Post, error) {

	post, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Only the author may update the post
	if post.AuthorID != authorID {
		return nil, errors.New("you are not allowed to update this post")
	}

	if req.Title != "" {
		post.Title = strings.TrimSpace(req.Title)
	}

	if req.Content != "" {
		post.Content = strings.TrimSpace(req.Content)
		post.IsEdited = true
	}

	if req.ContentType != "" {
		post.ContentType = strings.TrimSpace(req.ContentType)
	}

	if req.ImageURL != "" {
		post.ImageURL = req.ImageURL
	}

	if req.IsAnonymous != nil {
		post.IsAnonymous = *req.IsAnonymous
	}

	if req.IsPinned != nil {
		post.IsPinned = *req.IsPinned
	}

	if req.IsLocked != nil {
		post.IsLocked = *req.IsLocked
	}

	if err := s.repo.Update(post); err != nil {
		return nil, err
	}

	return post, nil
}

// ----------------------------------------------------
// Delete Post
// ----------------------------------------------------

func (s *service) DeletePost(id uint, authorID uint) error {

	post, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return err
	}

	// Only the author may delete the post
	if post.AuthorID != authorID {
		return errors.New("you are not allowed to delete this post")
	}

	return s.repo.Delete(id)
}
