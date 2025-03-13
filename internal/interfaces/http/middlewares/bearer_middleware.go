package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/BrunoPolaski/auth-service/internal/infra/thirdparty/jwt"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
)

func BearerMiddleware(jwtAdapter jwt.JWT) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encoder := json.NewEncoder(w)

			auth := r.Header.Get("Authorization")
			if !strings.Contains(auth, "Bearer ") {
				restErr := rest_err.NewBadRequestError("authentication is invalid")

				w.WriteHeader(restErr.Code)
				encoder.Encode(restErr)
				return
			}

			token := strings.Split(auth, " ")[1]

			parsedToken, err := jwtAdapter.ParseToken(token)
			if err != nil {
				w.WriteHeader(err.Code)
				encoder.Encode(err)
				return
			}

			exp, _ := parsedToken.Claims.GetExpirationTime()
			if exp.Before(time.Now()) {
				restErr := rest_err.NewUnauthorizedError("token expired")
				w.WriteHeader(restErr.Code)
				encoder.Encode(restErr)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
