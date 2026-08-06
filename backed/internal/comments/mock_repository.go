package comments

import "errors"

type MockRepository struct {
	comments map[uint]*Comment
	nextID   uint
}

func NewMockRepository() Repository {
	return &MockRepository{
		comments: make(map[uint]*Comment),
		nextID:   1,
	}
}

// Create a comment
func (m *MockRepository) Create(comment *Comment) error {
	comment.ID = m.nextID
	m.comments[m.nextID] = comment
	m.nextID++

	return nil
}

// Get a comment by ID
func (m *MockRepository) GetByID(id uint) (*Comment, error) {
	comment, ok := m.comments[id]
	if !ok {
		return nil, errors.New("comment not found")
	}

	return comment, nil
}

// Get all comments for a post
func (m *MockRepository) GetByPostID(postID uint) ([]Comment, error) {
	var comments []Comment

	for _, comment := range m.comments {
		if comment.PostID == postID {
			comments = append(comments, *comment)
		}
	}

	return comments, nil
}

// Update a comment
func (m *MockRepository) Update(comment *Comment) error {
	if _, ok := m.comments[comment.ID]; !ok {
		return errors.New("comment not found")
	}

	m.comments[comment.ID] = comment
	return nil
}

// Delete a comment
func (m *MockRepository) Delete(id uint) error {
	if _, ok := m.comments[id]; !ok {
		return errors.New("comment not found")
	}

	delete(m.comments, id)
	return nil
}
