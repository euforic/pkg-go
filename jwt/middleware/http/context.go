package jwthttp

import (
	"context"
	"errors"

	"github.com/euforic/pkg-go/jwt"
)

type jwtContextKey string

const ContextKey jwtContextKey = "jwt_context"

// TokenFromContext gets the raw token from the context and parses into a *Token
func TokenFromContext(ctx context.Context) (*jwt.Token, error) {
	tokenVal := ctx.Value(ContextKey)
	token, ok := tokenVal.(*jwt.Token)
	if !ok {
		return nil, errors.New("invalid or missing token")
	}
	if err := token.ParseError(); err != nil {
		return nil, err
	}

	return token, nil
}
