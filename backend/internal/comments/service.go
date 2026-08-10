package comments

import (
	"errors"
	"strings"
)

type Service interface {
	CreateComment(req *CreateCommentRequest, postID, authorID uint) (*Comment, error)
	GetComment(id uint) (*Comment, error)
	GetPostComments(postID uint) ([]Comment, error)
	UpdateComment(id, authorID uint, req *UpdateCommentRequest) (*Comment, error)
	DeleteComment(id, authorID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create a new comment
func (s *service) CreateComment(req *CreateCommentRequest, postID, authorID uint) (*Comment, error) {
	content := strings.TrimSpace(req.Content)

	if content == "" {
		return nil, errors.New("comment content is required")
	}

	comment := &Comment{
		PostID:   postID,
		AuthorID: authorID,
		Content:  content,
		IsEdited: false,
	}

	if err := s.repo.Create(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// Get a comment by ID
func (s *service) GetComment(id uint) (*Comment, error) {
	return s.repo.GetByID(id)
}

// Get all comments for a post
func (s *service) GetPostComments(postID uint) ([]Comment, error) {
	return s.repo.GetByPostID(postID)
}

// Update a comment
func (s *service) UpdateComment(id, authorID uint, req *UpdateCommentRequest) (*Comment, error) {
	comment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if comment.AuthorID != authorID {
		return nil, errors.New("you are not authorized to update this comment")
	}

	content := strings.TrimSpace(req.Content)
	if content == "" {
		return nil, errors.New("comment content is required")
	}

	comment.Content = content
	comment.IsEdited = true

	if err := s.repo.Update(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// Delete a comment
func (s *service) DeleteComment(id, authorID uint) error {
	comment, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if comment.AuthorID != authorID {
		return errors.New("you are not authorized to delete this comment")
	}

	return s.repo.Delete(id)
}
