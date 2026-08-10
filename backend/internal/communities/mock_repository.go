package communities

import (
	"gorm.io/gorm"
)

type MockRepository struct {
	communities map[uint]*Community
	nextID      uint
}

func NewMockRepository() Repository {
	return &MockRepository{
		communities: make(map[uint]*Community),
		nextID:      1,
	}
}

func (m *MockRepository) Create(community *Community) error {
	community.ID = m.nextID
	m.communities[m.nextID] = community
	m.nextID++

	return nil
}

func (m *MockRepository) GetByID(id uint) (*Community, error) {
	community, ok := m.communities[id]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}

	return community, nil
}

func (m *MockRepository) GetByName(name string) (*Community, error) {
	for _, community := range m.communities {
		if community.Name == name {
			return community, nil
		}
	}

	return nil, gorm.ErrRecordNotFound
}

func (m *MockRepository) GetAll() ([]Community, error) {
	communities := make([]Community, 0, len(m.communities))

	for _, community := range m.communities {
		communities = append(communities, *community)
	}

	return communities, nil
}

func (m *MockRepository) Update(community *Community) error {
	if _, ok := m.communities[community.ID]; !ok {
		return gorm.ErrRecordNotFound
	}

	m.communities[community.ID] = community
	return nil
}

func (m *MockRepository) Delete(id uint) error {
	if _, ok := m.communities[id]; !ok {
		return gorm.ErrRecordNotFound
	}

	delete(m.communities, id)
	return nil
}
