package communities

type CreateCommunityRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"max=1000"`
	Category    string `json:"category" binding:"required"`
	BannerURL   string `json:"banner_url"`
	IsPrivate   bool   `json:"is_private"`
}

type UpdateCommunityRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	BannerURL   string `json:"banner_url"`
	IsPrivate   *bool  `json:"is_private"`
}

type CommunityResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	CreatorID   uint   `json:"creator_id"`
	BannerURL   string `json:"banner_url"`
	IsPrivate   bool   `json:"is_private"`
}

func ToCommunityResponse(c *Community) CommunityResponse {
	return CommunityResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Category:    c.Category,
		CreatorID:   c.CreatorID,
		BannerURL:   c.BannerURL,
		IsPrivate:   c.IsPrivate,
	}
}

func ToCommunityResponses(communities []Community) []CommunityResponse {
	responses := make([]CommunityResponse, 0, len(communities))

	for _, community := range communities {
		responses = append(responses, ToCommunityResponse(&community))
	}

	return responses
}
