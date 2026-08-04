package search

import "time"

// SearchRequest represents a search request.
type SearchRequest struct {
	Query string     `json:"query" binding:"required"`
	Type  SearchType `json:"type"`
}

// SearchResponse represents a single search history entry.
type SearchResponse struct {
	ID          uint       `json:"id"`
	UserID      uint       `json:"user_id"`
	Query       string     `json:"query"`
	Type        SearchType `json:"type"`
	ResultCount int        `json:"result_count"`
	CreatedAt   time.Time  `json:"created_at"`
}

// SearchResult represents a generic search result.
type SearchResult struct {
	ID       uint       `json:"id"`
	Type     SearchType `json:"type"`
	Title    string     `json:"title"`
	Subtitle string     `json:"subtitle,omitempty"`
}

// GlobalSearchResponse represents the response returned from a global search.
type GlobalSearchResponse struct {
	Query   string         `json:"query"`
	Count   int            `json:"count"`
	Results []SearchResult `json:"results"`
}
