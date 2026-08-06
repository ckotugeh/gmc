package communities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCommunity(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := &CreateCommunityRequest{
		Name:        "Go Developers",
		Description: "Community for Go developers",
		Category:    "Technology",
		IsPrivate:   false,
	}

	community, err := service.CreateCommunity(req, 1)

	assert.NoError(t, err)
	assert.NotNil(t, community)
	assert.Equal(t, uint(1), community.ID)
	assert.Equal(t, "Go Developers", community.Name)
	assert.Equal(t, "Technology", community.Category)
	assert.Equal(t, uint(1), community.CreatorID)
}
