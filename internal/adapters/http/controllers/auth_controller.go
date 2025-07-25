package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrunoPolaski/auth-service/internal/adapters/http/dto"
	"github.com/BrunoPolaski/auth-service/internal/adapters/services"
	"github.com/BrunoPolaski/auth-service/internal/infra/logger"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(service services.AuthService) AuthController {
	return &authController{
		authService: service,
	}
}

func (ac *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	logger.Info("Authenticating user")
	encoder := json.NewEncoder(w)

	username, password, _ := r.BasicAuth()

	token, refreshToken, err := ac.authService.SignIn(username, password)
	if err != nil {
		logger.Error(err.Message)
		w.WriteHeader(err.Code)
		encoder.Encode(err)
		return
	}

	response := dto.SigninResponse{
		RefreshToken: refreshToken,
		AccessToken:  token,
	}

	encoder.Encode(response)
}

func (ac *authController) RefreshToken(w http.ResponseWriter, r *http.Request) {

}
