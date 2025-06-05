package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrunoPolaski/auth-service/internal/adapters/services"
	"github.com/BrunoPolaski/auth-service/internal/config/logger"
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
	encoder := json.NewEncoder(w)

	logger.Info("Authenticating user")
	username, password, ok := r.BasicAuth()
	if !ok {
		httpErr := rest_err.NewUnauthorizedError("Basic auth header not found")
		logger.Error(httpErr.Message)
		w.WriteHeader(httpErr.Code)
		encoder.Encode(httpErr)
		return
	}

	token, err := ac.authService.SignIn(username, password)
	if err != nil {
		logger.Error(err.Message)
		w.WriteHeader(err.Code)
		encoder.Encode(err)
		return
	}

	encoder.Encode(map[string]string{
		"token": token,
	})
}
