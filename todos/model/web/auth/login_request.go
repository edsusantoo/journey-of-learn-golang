package auth

type LoginRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
}
