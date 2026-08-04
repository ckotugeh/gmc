package search

import (
	"errors"
	"strings"
)

var (
	ErrSearchNotFound = errors.New("search not found")
	ErrInvalidSearch  = errors.New("invalid search")
)

// Service defines the search business logic.
type Service interface {
	Create(req SearchRequest, userID uint) (*Search, error)
	GetByID(id uint) (*Search, error)
	GetByUserID(userID uint) ([]Search, error)
	GetAll() ([]Search, error)
	Delete(id uint) error

	SearchDoctors(query string) ([]SearchResult, error)
	SearchPatients(query string) ([]SearchResult, error)
	SearchHospitals(query string) ([]SearchResult, error)
	SearchCommunities(query string) ([]SearchResult, error)
	SearchPosts(query string) ([]SearchResult, error)
	GlobalSearch(query string) ([]SearchResult, error)
}

type service struct {
	repo Repository
}

// NewService creates a new search service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create saves a search record.
func (s *service) Create(req SearchRequest, userID uint) (*Search, error) {
	query := strings.TrimSpace(req.Query)

	if userID == 0 || query == "" {
		return nil, ErrInvalidSearch
	}

	var (
		results []SearchResult
		err     error
	)

	switch req.Type {
	case TypeDoctor:
		results, err = s.repo.SearchDoctors(query)
	case TypePatient:
		results, err = s.repo.SearchPatients(query)
	case TypeHospital:
		results, err = s.repo.SearchHospitals(query)
	case TypeCommunity:
		results, err = s.repo.SearchCommunities(query)
	case TypePost:
		results, err = s.repo.SearchPosts(query)
	default:
		results, err = s.repo.GlobalSearch(query)
	}

	if err != nil {
		return nil, err
	}

	search := &Search{
		UserID:      userID,
		Query:       query,
		Type:        req.Type,
		ResultCount: len(results),
	}

	if err := s.repo.Create(search); err != nil {
		return nil, err
	}

	return search, nil
}

// GetByID retrieves a search record.
func (s *service) GetByID(id uint) (*Search, error) {
	search, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrSearchNotFound
	}

	return search, nil
}

// GetByUserID retrieves a user's search history.
func (s *service) GetByUserID(userID uint) ([]Search, error) {
	return s.repo.GetByUserID(userID)
}

// GetAll retrieves all search history.
func (s *service) GetAll() ([]Search, error) {
	return s.repo.GetAll()
}

// Delete removes a search record.
func (s *service) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return ErrSearchNotFound
	}

	return s.repo.Delete(id)
}

// SearchDoctors searches doctors.
func (s *service) SearchDoctors(query string) ([]SearchResult, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, ErrInvalidSearch
	}
	return s.repo.SearchDoctors(query)
}

// SearchPatients searches patients.
func (s *service) SearchPatients(query string) ([]SearchResult, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, ErrInvalidSearch
	}
	return s.repo.SearchPatients(query)
}

// SearchHospitals searches hospitals.
func (s *service) SearchHospitals(query string) ([]SearchResult, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, ErrInvalidSearch
	}
	return s.repo.SearchHospitals(query)
}

// SearchCommunities searches communities.
func (s *service) SearchCommunities(query string) ([]SearchResult, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, ErrInvalidSearch
	}
	return s.repo.SearchCommunities(query)
}

// SearchPosts searches posts.
func (s *service) SearchPosts(query string) ([]SearchResult, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, ErrInvalidSearch
	}
	return s.repo.SearchPosts(query)
}

// GlobalSearch searches all supported entities.
func (s *service) GlobalSearch(query string) ([]SearchResult, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, ErrInvalidSearch
	}

	return s.repo.GlobalSearch(query)
}
