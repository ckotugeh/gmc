package search

import (
	"strings"

	"gorm.io/gorm"
)

// Repository defines the search repository contract.
type Repository interface {
	Create(search *Search) error
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

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new search repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create stores a search history record.
func (r *repository) Create(search *Search) error {
	return r.db.Create(search).Error
}

// GetByID retrieves a search history record by ID.
func (r *repository) GetByID(id uint) (*Search, error) {
	var search Search
	if err := r.db.First(&search, id).Error; err != nil {
		return nil, err
	}
	return &search, nil
}

// GetByUserID retrieves all searches made by a user.
func (r *repository) GetByUserID(userID uint) ([]Search, error) {
	var searches []Search
	if err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&searches).Error; err != nil {
		return nil, err
	}
	return searches, nil
}

// GetAll retrieves all search history.
func (r *repository) GetAll() ([]Search, error) {
	var searches []Search
	if err := r.db.Order("created_at DESC").Find(&searches).Error; err != nil {
		return nil, err
	}
	return searches, nil
}

// Delete removes a search history record.
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Search{}, id).Error
}

// SearchDoctors searches doctors.
func (r *repository) SearchDoctors(query string) ([]SearchResult, error) {
	return r.searchByType(query, TypeDoctor)
}

// SearchPatients searches patients.
func (r *repository) SearchPatients(query string) ([]SearchResult, error) {
	return r.searchByType(query, TypePatient)
}

// SearchHospitals searches hospitals.
func (r *repository) SearchHospitals(query string) ([]SearchResult, error) {
	return r.searchByType(query, TypeHospital)
}

// SearchCommunities searches communities.
func (r *repository) SearchCommunities(query string) ([]SearchResult, error) {
	return r.searchByType(query, TypeCommunity)
}

// SearchPosts searches posts.
func (r *repository) SearchPosts(query string) ([]SearchResult, error) {
	return r.searchByType(query, TypePost)
}

// GlobalSearch searches across all supported types.
func (r *repository) GlobalSearch(query string) ([]SearchResult, error) {
	var results []SearchResult

	searches, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)

	for _, search := range searches {
		if strings.Contains(strings.ToLower(search.Query), query) {
			results = append(results, SearchResult{
				ID:       search.ID,
				Type:     search.Type,
				Title:    search.Query,
				Subtitle: "Search history",
			})
		}
	}

	return results, nil
}

// searchByType searches history by type.
func (r *repository) searchByType(query string, searchType SearchType) ([]SearchResult, error) {
	var results []SearchResult

	searches, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)

	for _, search := range searches {
		if search.Type == searchType &&
			strings.Contains(strings.ToLower(search.Query), query) {

			results = append(results, SearchResult{
				ID:       search.ID,
				Type:     search.Type,
				Title:    search.Query,
				Subtitle: "Search history",
			})
		}
	}

	return results, nil
}
