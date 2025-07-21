package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt_adapter "github.com/BrunoPolaski/auth-service/internal/infra/jwt"
	"github.com/BrunoPolaski/auth-service/internal/infra/logger"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
	"github.com/golang-jwt/jwt/v5"
)

func BearerAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		header := r.Header.Get("Authorization")
		jwtAdapter := jwt_adapter.NewJWTAdapter()

		token, restErr := jwtAdapter.TrimPrefix(header)
		if restErr != nil {
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		parsedToken, restErr := jwtAdapter.ParseToken(token)
		if restErr != nil {
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			logger.Error("Invalid token claims or token is not valid")

			restErr := rest_err.NewUnauthorizedError("invalid token")
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		exp, err := claims.GetExpirationTime()
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to get expiration time from claims: %s", err.Error()))

			restErr := rest_err.NewUnauthorizedError("invalid token")
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		if exp.Before(time.Now()) {
			restErr := rest_err.NewUnauthorizedError("token expired")
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		next.ServeHTTP(w, r)
	})
}
