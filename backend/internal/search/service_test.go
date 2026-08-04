package search

import "testing"

func TestCreateSearch(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := SearchRequest{
		Query: "cardiologist",
		Type:  TypeDoctor,
	}

	search, err := service.Create(req, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if search.ID == 0 {
		t.Fatal("expected search ID to be assigned")
	}

	if search.Query != req.Query {
		t.Fatalf("expected query %q, got %q", req.Query, search.Query)
	}

	if search.UserID != 1 {
		t.Fatalf("expected user ID 1, got %d", search.UserID)
	}
}

func TestCreateSearchValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []struct {
		name   string
		req    SearchRequest
		userID uint
	}{
		{
			name:   "empty query",
			req:    SearchRequest{},
			userID: 1,
		},
		{
			name: "blank query",
			req: SearchRequest{
				Query: "   ",
			},
			userID: 1,
		},
		{
			name: "invalid user",
			req: SearchRequest{
				Query: "doctor",
			},
			userID: 0,
		},
	}

	for _, tt := range tests {
		_, err := service.Create(tt.req, tt.userID)
		if err != ErrInvalidSearch {
			t.Fatalf("%s: expected %v, got %v", tt.name, ErrInvalidSearch, err)
		}
	}
}

func TestGetSearch(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(SearchRequest{
		Query: "neurologist",
		Type:  TypeDoctor,
	}, 1)

	search, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if search.ID != created.ID {
		t.Fatalf("expected ID %d, got %d", created.ID, search.ID)
	}
}

func TestGetUserSearches(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 3; i++ {
		_, _ = service.Create(SearchRequest{
			Query: "doctor",
			Type:  TypeDoctor,
		}, 1)
	}

	searches, err := service.GetByUserID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(searches) != 3 {
		t.Fatalf("expected 3 searches, got %d", len(searches))
	}
}

func TestGetAllSearches(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 5; i++ {
		_, _ = service.Create(SearchRequest{
			Query: "hospital",
			Type:  TypeHospital,
		}, uint(i+1))
	}

	searches, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(searches) != 5 {
		t.Fatalf("expected 5 searches, got %d", len(searches))
	}
}

func TestDeleteSearch(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(SearchRequest{
		Query: "community",
		Type:  TypeCommunity,
	}, 1)

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrSearchNotFound {
		t.Fatalf("expected %v, got %v", ErrSearchNotFound, err)
	}
}

func TestSearchDoctors(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(SearchRequest{
		Query: "cardiologist",
		Type:  TypeDoctor,
	}, 1)

	results, err := service.SearchDoctors("cardiologist")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestSearchHospitals(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(SearchRequest{
		Query: "Nairobi Hospital",
		Type:  TypeHospital,
	}, 1)

	results, err := service.SearchHospitals("Nairobi Hospital")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestGlobalSearch(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(SearchRequest{
		Query: "heart",
		Type:  TypeDoctor,
	}, 1)

	_, _ = service.Create(SearchRequest{
		Query: "heart",
		Type:  TypeHospital,
	}, 1)

	results, err := service.GlobalSearch("heart")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
}
