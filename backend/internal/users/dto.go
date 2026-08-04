package users

// CreateUserRequest represents the payload for creating a user.
type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}

// UpdateUserRequest represents the payload for updating a user.
type UpdateUserRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Role      *string `json:"role,omitempty"`
	IsActive  *bool   `json:"is_active,omitempty"`
}

// UserResponse represents a user returned to the client.
// Notice that Password is intentionally omitted.
type UserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	IsActive  bool   `json:"is_active"`
}

// ToResponse converts a User model into a response DTO.
func ToResponse(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		IsActive:  user.IsActive,
	}
}

// ToResponseList converts a slice of users into response DTOs.
func ToResponseList(users []User) []UserResponse {
	response := make([]UserResponse, 0, len(users))

	for _, user := range users {
		response = append(response, ToResponse(&user))
	}

	return response
}
