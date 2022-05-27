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
		if len(authHeader) < 2 {
			next.ServeHTTP(w, r)
			return
		}

		token, _ := j.Parse(authHeader[1], true)

		ctx := context.WithValue(context.Background(), ContextKey, token)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
