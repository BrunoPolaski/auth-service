package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/BrunoPolaski/auth-service/internal/adapters/repositories"
	"github.com/BrunoPolaski/auth-service/internal/infra/database"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

func ApiKeyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		apiKey := r.Header.Get("X-Api-Key")
		if strings.TrimSpace(apiKey) == "" {
			restErr := rest_err.NewUnauthorizedError("API key header not found or invalid")
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		db, _ := database.NewPostgresAdapter().Connect()

		apiKeyRepository := repositories.NewApiKeyRepository(db)

		_, err := apiKeyRepository.GetById(apiKey)
		if err != nil {
			restErr := rest_err.NewUnauthorizedError("Invalid API key")
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		next.ServeHTTP(w, r)
	})
}
