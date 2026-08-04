package posts

import (
	"errors"
)

type MockRepository struct {
	posts  map[uint]*Post
	nextID uint
}

func NewMockRepository() Repository {
	return &MockRepository{
		posts:  make(map[uint]*Post),
		nextID: 1,
	}
}

// Create
func (m *MockRepository) Create(post *Post) error {
	post.ID = m.nextID
	m.posts[m.nextID] = post
	m.nextID++

	return nil
}

// Get by ID
func (m *MockRepository) GetByID(id uint) (*Post, error) {

	post, ok := m.posts[id]
	if !ok {
		return nil, errors.New("post not found")
	}

	return post, nil
}

// Get all posts in a community
func (m *MockRepository) GetByCommunityID(communityID uint) ([]Post, error) {

	var posts []Post

	for _, post := range m.posts {
		if post.CommunityID == communityID {
			posts = append(posts, *post)
		}
	}

	return posts, nil
}

// Update
func (m *MockRepository) Update(post *Post) error {

	if _, ok := m.posts[post.ID]; !ok {
		return errors.New("post not found")
	}

	m.posts[post.ID] = post

	return nil
}

// Delete
func (m *MockRepository) Delete(id uint) error {

	if _, ok := m.posts[id]; !ok {
		return errors.New("post not found")
	}

	delete(m.posts, id)

	return nil
}
