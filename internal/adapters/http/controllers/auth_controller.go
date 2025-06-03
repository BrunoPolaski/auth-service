package controller

import (
	"encoding/json"
	"net/http"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/domain/service"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
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
