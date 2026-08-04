package communities

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Service interface {
	CreateCommunity(req *CreateCommunityRequest, creatorID uint) (*Community, error)
	GetCommunity(id uint) (*Community, error)
	GetCommunities() ([]Community, error)
	UpdateCommunity(id uint, req *UpdateCommunityRequest) (*Community, error)
	DeleteCommunity(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateCommunity(req *CreateCommunityRequest, creatorID uint) (*Community, error) {
	name := strings.TrimSpace(req.Name)
	category := strings.TrimSpace(req.Category)

	if name == "" {
		return nil, errors.New("community name is required")
	}

	if category == "" {
		return nil, errors.New("category is required")
	}

	existing, err := s.repo.GetByName(name)
	if err == nil && existing != nil {
		return nil, errors.New("community already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	community := &Community{
		Name:        name,
		Description: strings.TrimSpace(req.Description),
		Category:    category,
		CreatorID:   creatorID,
		BannerURL:   strings.TrimSpace(req.BannerURL),
		IsPrivate:   req.IsPrivate,
	}

	if err := s.repo.Create(community); err != nil {
		return nil, err
	}

	return community, nil
}

func (s *service) GetCommunity(id uint) (*Community, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetCommunities() ([]Community, error) {
	return s.repo.GetAll()
}

func (s *service) UpdateCommunity(id uint, req *UpdateCommunityRequest) (*Community, error) {
	community, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		name := strings.TrimSpace(req.Name)

		if name != community.Name {
			existing, err := s.repo.GetByName(name)
			if err == nil && existing != nil && existing.ID != community.ID {
				return nil, errors.New("community name already exists")
			}

			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}

		community.Name = name
	}

	if req.Description != "" {
		community.Description = strings.TrimSpace(req.Description)
	}

	if req.Category != "" {
		community.Category = strings.TrimSpace(req.Category)
	}

	if req.BannerURL != "" {
		community.BannerURL = strings.TrimSpace(req.BannerURL)
	}

	if req.IsPrivate != nil {
		community.IsPrivate = *req.IsPrivate
	}

	if err := s.repo.Update(community); err != nil {
		return nil, err
	}

	return community, nil
}

func (s *service) DeleteCommunity(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
