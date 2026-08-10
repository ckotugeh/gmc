package reactions

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	CreateReaction(postID, userID uint, req *CreateReactionRequest) (*Reaction, error)
	GetPostReactions(postID uint) ([]Reaction, error)
	UpdateReaction(id, userID uint, req *UpdateReactionRequest) (*Reaction, error)
	DeleteReaction(id, userID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateReaction(postID, userID uint, req *CreateReactionRequest) (*Reaction, error) {
	// Check if the user has already reacted to this post
	existing, err := s.repo.GetByPostAndUser(postID, userID)

	if err == nil && existing != nil {
		// Update the existing reaction instead of creating another
		existing.ReactionType = req.ReactionType

		if err := s.repo.Update(existing); err != nil {
			return nil, err
		}

		return existing, nil
	}

	// Ignore "not found" and create a new reaction
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) &&
		err.Error() != "reaction not found" {
		return nil, err
	}

	reaction := &Reaction{
		PostID:       postID,
		UserID:       userID,
		ReactionType: req.ReactionType,
	}

	if err := s.repo.Create(reaction); err != nil {
		return nil, err
	}

	return reaction, nil
}

func (s *service) GetPostReactions(postID uint) ([]Reaction, error) {
	return s.repo.GetByPost(postID)
}

func (s *service) UpdateReaction(id, userID uint, req *UpdateReactionRequest) (*Reaction, error) {
	reaction, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if reaction.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	reaction.ReactionType = req.ReactionType

	if err := s.repo.Update(reaction); err != nil {
		return nil, err
	}

	return reaction, nil
}

func (s *service) DeleteReaction(id, userID uint) error {
	reaction, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if reaction.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(id)
}
