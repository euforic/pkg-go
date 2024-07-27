package jwthttp

import (
	"context"
	"errors"

	"github.com/euforic/pkg-go/jwt"
)

// ErrMissingToken is returned when the token is missing
var ErrMissingToken = errors.New("missing token")

// jwtContextKey is the key type for the context
type jwtContextKey string

// ContextKey is the key to use when setting the token in the context
const ContextKey jwtContextKey = "jwt_context"

// TokenFromContext gets the raw token from the context and parses into a *Token
func TokenFromContext(ctx context.Context) (*jwt.Token, error) {
	tokenVal := ctx.Value(ContextKey)
	token, ok := tokenVal.(*jwt.Token)
	if !ok {
		return nil, ErrMissingToken
	}

	return token, nil
}
