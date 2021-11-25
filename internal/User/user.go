package User

// CreateUserRequest represent struct for create request
type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

// UpdateUserRequest represent struct for update request
type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}
