package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateStack(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			curMiddleware := middlewares[i]
			next = curMiddleware(next)
		}

		return next
	}
}
