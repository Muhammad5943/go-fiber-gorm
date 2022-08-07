package request

type LoginRequest struct {
	Email    string `json:"email" validate:"required,omitempty,email"`
	Password string `json:"password" validate:"required"`
}
