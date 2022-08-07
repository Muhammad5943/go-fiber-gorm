package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,omitempty,email"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required,min=3,max=13"`
}

type UserUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone" validate:"required,min=3,max=13"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required,omitempty,email"`
}
