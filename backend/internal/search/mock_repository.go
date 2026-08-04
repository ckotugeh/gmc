package search

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	searches []Search
	nextID   uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		searches: []Search{},
		nextID:   1,
	}
}

// Create stores a search record.
func (m *MockRepository) Create(search *Search) error {
	search.ID = m.nextID
	m.nextID++
	m.searches = append(m.searches, *search)
	return nil
}

// GetByID retrieves a search by ID.
func (m *MockRepository) GetByID(id uint) (*Search, error) {
	for _, search := range m.searches {
		if search.ID == id {
			s := search
			return &s, nil
		}
	}
	return nil, ErrSearchNotFound
}

// GetByUserID retrieves searches by user ID.
func (m *MockRepository) GetByUserID(userID uint) ([]Search, error) {
	var results []Search

	for _, search := range m.searches {
		if search.UserID == userID {
			results = append(results, search)
		}
	}

	return results, nil
}

// GetAll retrieves all searches.
func (m *MockRepository) GetAll() ([]Search, error) {
	return m.searches, nil
}

// Delete deletes a search.
func (m *MockRepository) Delete(id uint) error {
	for i, search := range m.searches {
		if search.ID == id {
			m.searches = append(m.searches[:i], m.searches[i+1:]...)
			return nil
		}
	}
	return ErrSearchNotFound
}

// SearchDoctors searches doctors.
func (m *MockRepository) SearchDoctors(query string) ([]SearchResult, error) {
	return m.searchByType(query, TypeDoctor), nil
}

// SearchPatients searches patients.
func (m *MockRepository) SearchPatients(query string) ([]SearchResult, error) {
	return m.searchByType(query, TypePatient), nil
}

// SearchHospitals searches hospitals.
func (m *MockRepository) SearchHospitals(query string) ([]SearchResult, error) {
	return m.searchByType(query, TypeHospital), nil
}

// SearchCommunities searches communities.
func (m *MockRepository) SearchCommunities(query string) ([]SearchResult, error) {
	return m.searchByType(query, TypeCommunity), nil
}

// SearchPosts searches posts.
func (m *MockRepository) SearchPosts(query string) ([]SearchResult, error) {
	return m.searchByType(query, TypePost), nil
}

// GlobalSearch searches all records.
func (m *MockRepository) GlobalSearch(query string) ([]SearchResult, error) {
	var results []SearchResult

	results = append(results, m.searchByType(query, TypeDoctor)...)
	results = append(results, m.searchByType(query, TypePatient)...)
	results = append(results, m.searchByType(query, TypeHospital)...)
	results = append(results, m.searchByType(query, TypeCommunity)...)
	results = append(results, m.searchByType(query, TypePost)...)

	return results, nil
}

// searchByType performs an in-memory search.
func (m *MockRepository) searchByType(query string, searchType SearchType) []SearchResult {
	var results []SearchResult

	for _, search := range m.searches {
		if search.Type == searchType && search.Query == query {
			results = append(results, SearchResult{
				ID:       search.ID,
				Type:     search.Type,
				Title:    search.Query,
				Subtitle: "Search history",
			})
		}
	}

	return results
}
