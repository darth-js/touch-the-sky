package auth

import (
	"net/http"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// JWT verification logic
		next.ServeHTTP(w, r)
	})
}
