package services

import (
	"github.com/BrunoPolaski/auth-service/internal/adapters/repositories"
	valueobjects "github.com/BrunoPolaski/auth-service/internal/core/value_objects"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

type AuthService interface {
	SignIn(username, password string) (string, string, *rest_err.RestErr)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (as *authService) SignIn(username, password string) (string, string, *rest_err.RestErr) {
	pwd, err := valueobjects.NewPassword(password)
	if err != nil {
		return "", "", rest_err.NewBadRequestError(err.Error())
	}

	user, restErr := as.userRepository.FindUserByEmail(username)
	if restErr != nil {
		return "", "", restErr
	}

	err = pwd.ComparePassword(user.Password().Value())
	if err != nil {
		return "", "", rest_err.NewUnauthorizedError(err.Error())
	}

	return "token", "refreshToken", nil
}

func (as *authService) RefreshToken() {

}
