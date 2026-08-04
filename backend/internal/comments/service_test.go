package comments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreateCommentRequest{
		Content: "This is my first comment.",
	}

	comment, err := service.CreateComment(req, 1, 1)

	assert.NoError(t, err)
	assert.NotNil(t, comment)
	assert.Equal(t, uint(1), comment.ID)
	assert.Equal(t, uint(1), comment.PostID)
	assert.Equal(t, uint(1), comment.AuthorID)
	assert.Equal(t, "This is my first comment.", comment.Content)
	assert.False(t, comment.IsEdited)
}

func TestCreateCommentWithoutContent(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreateCommentRequest{
		Content: "",
	}

	comment, err := service.CreateComment(req, 1, 1)

	assert.Error(t, err)
	assert.Nil(t, comment)
}

func TestGetComment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.CreateComment(&CreateCommentRequest{
		Content: "Testing retrieval",
	}, 1, 1)

	comment, err := service.GetComment(created.ID)

	assert.NoError(t, err)
	assert.NotNil(t, comment)
	assert.Equal(t, created.ID, comment.ID)
	assert.Equal(t, "Testing retrieval", comment.Content)
}

func TestUpdateComment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.CreateComment(&CreateCommentRequest{
		Content: "Original comment",
	}, 1, 1)

	req := &UpdateCommentRequest{
		Content: "Updated comment",
	}

	updated, err := service.UpdateComment(created.ID, 1, req)

	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, "Updated comment", updated.Content)
	assert.True(t, updated.IsEdited)
}

func TestUpdateCommentUnauthorized(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.CreateComment(&CreateCommentRequest{
		Content: "Original comment",
	}, 1, 1)

	req := &UpdateCommentRequest{
		Content: "Hacked comment",
	}

	updated, err := service.UpdateComment(created.ID, 2, req)

	assert.Error(t, err)
	assert.Nil(t, updated)
}

func TestDeleteComment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.CreateComment(&CreateCommentRequest{
		Content: "Delete me",
	}, 1, 1)

	err := service.DeleteComment(created.ID, 1)

	assert.NoError(t, err)

	comment, err := service.GetComment(created.ID)
	assert.Error(t, err)
	assert.Nil(t, comment)
}

func TestDeleteCommentUnauthorized(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.CreateComment(&CreateCommentRequest{
		Content: "Protected comment",
	}, 1, 1)

	err := service.DeleteComment(created.ID, 2)

	assert.Error(t, err)
}
