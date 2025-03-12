package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateStack(middlewares []Middleware, finalHandler http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		finalHandler = middlewares[i](finalHandler)
	}
	return finalHandler
}
