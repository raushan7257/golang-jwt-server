package view

type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewUserResponse(id uint, email, firstName, lastName string) UserResponse {
	return UserResponse{
		ID:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}
