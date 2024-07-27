package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	// ErrTokenParse is returned when the token could not be parsed
	ErrTokenParse = errors.New("error parsing token")
	// ErrTokenValidate is returned when the token could not be validated
	ErrTokenValidate = errors.New("error validating token")
	// ErrTokenSign is returned when the token could not be signed
	ErrTokenSign = errors.New("error signing token")
	// ErrTokenCreate is returned when the token could not be created
	ErrTokenCreate = errors.New("error creating token")
	// ErrTokenExpired is returned when the token is expired
	ErrTokenExpired = errors.New("token expired")
	// ErrTokenInvalid is returned when the token is invalid
	ErrTokenInvalid = errors.New("token invalid")
)

// JWT is a struct that holds the private and public keys
type JWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// New creates a new instance of JWT util
func New(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT {
	return &JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// CreateAndSign a new jwt token
func (j JWT) CreatAndSign(ttl time.Duration, claims jwt.MapClaims) (string, error) {
	t, err := j.Create(ttl, claims)
	if err != nil {
		return "", err
	}

	return j.Sign(t)
}

// Create generates a new jwt token string
func (j JWT) Create(ttl time.Duration, claims jwt.MapClaims) (*Token, error) {
	now := time.Now().UTC()

	claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()          // The time at which the token was issued.

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return &Token{Token: token}, nil
}

// Create generates a new jwt token string
func (j JWT) Sign(token *Token) (string, error) {
	tokenStr, err := token.SignedString(j.privateKey)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return tokenStr, nil
}

// Parse takes in a jwt token string parses it, validates it and return a *Token
func (j JWT) Parse(token string, validate bool) (*Token, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method %s: %w", jwtToken.Header["alg"], ErrTokenParse)
		}

		var pubKey interface{} = j.publicKey
		if j.publicKey == nil || !validate {
			pubKey = []byte{}
		}

		return pubKey, nil
	})

	if tok == nil {
		return nil, fmt.Errorf("parse: %w", ErrTokenParse)
	}

	t := Token{
		Token: tok,
	}

	if err != nil && !validate && err.Error() != "key is of invalid type" {
		return nil, fmt.Errorf("parse: %w", err)
	}

	if !tok.Valid && validate {
		return nil, fmt.Errorf("parse: %w", ErrTokenValidate)
	}

	return &t, nil
}
