package jwthttp

import (
	"context"
	"net/http"
	"strings"

	"github.com/euforic/pkg-go/jwt"
)

// TokenMiddleware ...
func TokenMiddleware(j *jwt.JWT, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), " ")

		// Check if the token is present
		tokenPartsLen := 2
		if len(authHeader) != tokenPartsLen {
			next.ServeHTTP(w, r)

			return
		}

		// Parse the token
		token, err := j.Parse(authHeader[1], true)
		if err != nil {
			next.ServeHTTP(w, r)

			return
		}

		// Add the token to the context
		r = r.WithContext(context.WithValue(r.Context(), ContextKey, token))

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
