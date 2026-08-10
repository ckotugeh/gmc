package reactions

import "errors"

type MockRepository struct {
	reactions map[uint]*Reaction
	nextID    uint
}

func NewMockRepository() Repository {
	return &MockRepository{
		reactions: make(map[uint]*Reaction),
		nextID:    1,
	}
}

func (m *MockRepository) Create(reaction *Reaction) error {
	reaction.ID = m.nextID
	m.reactions[m.nextID] = reaction
	m.nextID++

	return nil
}

func (m *MockRepository) GetByID(id uint) (*Reaction, error) {
	reaction, ok := m.reactions[id]
	if !ok {
		return nil, errors.New("reaction not found")
	}

	return reaction, nil
}

func (m *MockRepository) GetByPostAndUser(postID, userID uint) (*Reaction, error) {
	for _, reaction := range m.reactions {
		if reaction.PostID == postID && reaction.UserID == userID {
			return reaction, nil
		}
	}

	return nil, errors.New("reaction not found")
}

func (m *MockRepository) GetByPost(postID uint) ([]Reaction, error) {
	var reactions []Reaction

	for _, reaction := range m.reactions {
		if reaction.PostID == postID {
			reactions = append(reactions, *reaction)
		}
	}

	return reactions, nil
}

func (m *MockRepository) Update(reaction *Reaction) error {
	if _, ok := m.reactions[reaction.ID]; !ok {
		return errors.New("reaction not found")
	}

	m.reactions[reaction.ID] = reaction
	return nil
}

func (m *MockRepository) Delete(id uint) error {
	if _, ok := m.reactions[id]; !ok {
		return errors.New("reaction not found")
	}

	delete(m.reactions, id)
	return nil
}
