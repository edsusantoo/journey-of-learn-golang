package auth

type RegisterRequest struct {
	Username  string `validate:"required,min=6" json:"username"`
	FirstName string `validate:"required" json:"firstname"`
	LastName  string `validate:"required" json:"lastname"`
	Password  string `validate:"required" json:"password"`
}
