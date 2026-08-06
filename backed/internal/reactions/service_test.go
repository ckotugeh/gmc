package reactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReaction(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreateReactionRequest{
		ReactionType: ReactionLike,
	}

	reaction, err := service.CreateReaction(1, 1, req)

	assert.NoError(t, err)
	assert.NotNil(t, reaction)
	assert.Equal(t, uint(1), reaction.ID)
	assert.Equal(t, uint(1), reaction.PostID)
	assert.Equal(t, uint(1), reaction.UserID)
	assert.Equal(t, ReactionLike, reaction.ReactionType)
}

func TestCreateReactionUpdatesExisting(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, err := service.CreateReaction(1, 1, &CreateReactionRequest{
		ReactionType: ReactionLike,
	})
	assert.NoError(t, err)

	reaction, err := service.CreateReaction(1, 1, &CreateReactionRequest{
		ReactionType: ReactionDislike,
	})

	assert.NoError(t, err)
	assert.NotNil(t, reaction)
	assert.Equal(t, uint(1), reaction.ID)
	assert.Equal(t, ReactionDislike, reaction.ReactionType)
}

func TestGetPostReactions(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreateReaction(1, 1, &CreateReactionRequest{
		ReactionType: ReactionLike,
	})

	_, _ = service.CreateReaction(1, 2, &CreateReactionRequest{
		ReactionType: ReactionDislike,
	})

	reactions, err := service.GetPostReactions(1)

	assert.NoError(t, err)
	assert.Len(t, reactions, 2)
}

func TestUpdateReaction(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	reaction, _ := service.CreateReaction(1, 1, &CreateReactionRequest{
		ReactionType: ReactionLike,
	})

	updated, err := service.UpdateReaction(reaction.ID, 1, &UpdateReactionRequest{
		ReactionType: ReactionDislike,
	})

	assert.NoError(t, err)
	assert.Equal(t, ReactionDislike, updated.ReactionType)
}

func TestUpdateReactionUnauthorized(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	reaction, _ := service.CreateReaction(1, 1, &CreateReactionRequest{
		ReactionType: ReactionLike,
	})

	updated, err := service.UpdateReaction(reaction.ID, 2, &UpdateReactionRequest{
		ReactionType: ReactionDislike,
	})

	assert.Error(t, err)
	assert.Nil(t, updated)
	assert.Equal(t, "unauthorized", err.Error())
}

func TestDeleteReaction(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	reaction, _ := service.CreateReaction(1, 1, &CreateReactionRequest{
		ReactionType: ReactionLike,
	})

	err := service.DeleteReaction(reaction.ID, 1)

	assert.NoError(t, err)
}

func TestDeleteReactionUnauthorized(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	reaction, _ := service.CreateReaction(1, 1, &CreateReactionRequest{
		ReactionType: ReactionLike,
	})

	err := service.DeleteReaction(reaction.ID, 2)

	assert.Error(t, err)
	assert.Equal(t, "unauthorized", err.Error())
}
