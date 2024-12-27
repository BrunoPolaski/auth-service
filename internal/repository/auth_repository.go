package repository

type AuthRepository interface {
	SignIn(username, password string) error
}

type authRepository struct {
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (ar authRepository) SignIn(username, password string) error {
	return nil
}
