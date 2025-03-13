package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrunoPolaski/auth-service/internal/application/services"
	"github.com/BrunoPolaski/auth-service/internal/config/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	ForgotPassword(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
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

	logger.Info("User authenticated")
}

func (ac *authController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	logger.Info("Refreshing token")
	token := r.Header.Get("Authorization")
	newToken, err := ac.authService.RefreshToken(token)
	if err != nil {
		logger.Error(err.Message)
		w.WriteHeader(err.Code)
		encoder.Encode(err)
		return
	}

	encoder.Encode(map[string]string{
		"token": newToken,
	})

	logger.Info("Token refreshed")
}

func (ac *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	logger.Info("Creating user")
}
