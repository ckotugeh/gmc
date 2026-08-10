package posts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreatePostRequest{
		CommunityID: 1,
		Title:       "Heart Failure Guidelines",
		Content:     "Let's discuss the latest ESC guidelines.",
		ContentType: "markdown",
	}

	post, err := service.CreatePost(req, 1)

	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, uint(1), post.ID)
	assert.Equal(t, uint(1), post.CommunityID)
	assert.Equal(t, uint(1), post.AuthorID)
	assert.Equal(t, "Heart Failure Guidelines", post.Title)
	assert.Equal(t, "Let's discuss the latest ESC guidelines.", post.Content)
	assert.Equal(t, "markdown", post.ContentType)
}

func TestCreatePostWithoutTitle(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreatePostRequest{
		CommunityID: 1,
		Content:     "Some content",
	}

	post, err := service.CreatePost(req, 1)

	assert.Error(t, err)
	assert.Nil(t, post)
	assert.Equal(t, "title is required", err.Error())
}

func TestCreatePostWithoutContent(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreatePostRequest{
		CommunityID: 1,
		Title:       "Test Post",
	}

	post, err := service.CreatePost(req, 1)

	assert.Error(t, err)
	assert.Nil(t, post)
	assert.Equal(t, "content is required", err.Error())
}

func TestGetPost(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreatePostRequest{
		CommunityID: 1,
		Title:       "Cardiology Discussion",
		Content:     "Discussion content",
	}

	created, _ := service.CreatePost(req, 1)

	post, err := service.GetPost(created.ID)

	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, created.ID, post.ID)
	assert.Equal(t, created.Title, post.Title)
}

func TestUpdatePost(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreatePostRequest{
		CommunityID: 1,
		Title:       "Original Title",
		Content:     "Original Content",
	}

	post, _ := service.CreatePost(req, 1)

	update := &UpdatePostRequest{
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	updated, err := service.UpdatePost(post.ID, 1, update)

	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, "Updated Title", updated.Title)
	assert.Equal(t, "Updated Content", updated.Content)
	assert.True(t, updated.IsEdited)
}

func TestUpdatePostUnauthorized(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreatePostRequest{
		CommunityID: 1,
		Title:       "Original",
		Content:     "Content",
	}

	post, _ := service.CreatePost(req, 1)

	update := &UpdatePostRequest{
		Title: "New Title",
	}

	updated, err := service.UpdatePost(post.ID, 2, update)

	assert.Error(t, err)
	assert.Nil(t, updated)
	assert.Equal(t, "you are not allowed to update this post", err.Error())
}

func TestDeletePost(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreatePostRequest{
		CommunityID: 1,
		Title:       "Delete Me",
		Content:     "Content",
	}

	post, _ := service.CreatePost(req, 1)

	err := service.DeletePost(post.ID, 1)

	assert.NoError(t, err)

	_, err = service.GetPost(post.ID)
	assert.Error(t, err)
}

func TestDeletePostUnauthorized(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreatePostRequest{
		CommunityID: 1,
		Title:       "Protected",
		Content:     "Content",
	}

	post, _ := service.CreatePost(req, 1)

	err := service.DeletePost(post.ID, 2)

	assert.Error(t, err)
	assert.Equal(t, "you are not allowed to delete this post", err.Error())
}
