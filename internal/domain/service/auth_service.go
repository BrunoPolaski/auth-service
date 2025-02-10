package service

type AuthService interface {
	SignIn(username string, password string) (bool, error)
}

type authService struct {
	authRepository AuthRepository
}

func NewAuthService(authRepository AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}
