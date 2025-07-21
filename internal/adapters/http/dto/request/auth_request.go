package request

type SigninRequest struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}
