package jwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// Token represents a JWT token
type Token struct {
	*jwt.Token
}

// NewToken creates a new instance of Token from a *jwt.Token
func NewToken(t *jwt.Token) *Token {
	return &Token{Token: t}
}

func (t Token) Claims() (jwt.MapClaims, bool) {
	value, ok := t.Token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}

	return value, true
}

func (t Token) Get(key string) (interface{}, bool) {
	v, ok := get[any](t, key)
	if !ok {
		return nil, false
	}

	return v, true
}

// GetString returns the value associated with the key as a string.
func (t Token) GetString(key string) string {
	v, _ := get[string](t, key)

	return v
}

// GetBool returns the value associated with the key as a boolean.
func (t Token) GetBool(key string) bool {
	v, _ := get[bool](t, key)

	return v
}

// GetInt returns the value associated with the key as an integer.
func (t Token) GetInt(key string) int {
	v, _ := get[int](t, key)

	return v
}

// GetUint returns the value associated with the key as an unsigned integer.
func (t Token) GetUint(key string) uint {
	v, _ := get[uint](t, key)

	return v
}

// GetFloat32 returns the value associated with the key as a float32.
func (t Token) GetFloat32(key string) float32 {
	v, _ := get[float32](t, key)

	return v
}

// GetFloat64 returns the value associated with the key as a float64.
func (t Token) GetFloat64(key string) float64 {
	v, _ := get[float64](t, key)

	return v
}

// GetTime returns the value associated with the key as time.
func (t Token) GetTime(key string) time.Time {
	v, _ := get[time.Time](t, key)

	return v
}

// GetDuration returns the value associated with the key as a duration.
func (t Token) GetDuration(key string) time.Duration {
	v, _ := get[time.Duration](t, key)

	return v
}

// GetSlice returns the value associated with the key as a slice of interface{}.
func (t Token) GetSlice(key string) []interface{} {
	v, ok := get[[]interface{}](t, key)
	if !ok {
		return nil
	}

	return v
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (t Token) GetStringSlice(key string) []string {
	v, ok := get[[]string](t, key)
	if !ok {
		return nil
	}

	return v
}

// tokenVal is a type that can be used to represent any value that can be stored in a token.
type tokenVal interface {
	int | uint | string | bool | float32 | float64 | time.Time | time.Duration | []any | []string | any
}

// get retrieves the value associated with the key from the token.
func get[T tokenVal](t Token, key string) (T, bool) { //nolint:ireturn
	var value interface{}
	var exists bool

	keys := strings.Split(key, ".")
	for i, key := range keys {
		if i == 0 {
			value, exists = t.Token.Claims.(jwt.MapClaims)[keys[0]]
			if !exists {
				return *new(T), false
			}

			continue
		}
		if value == nil {
			return *new(T), false
		}
		value, exists = value.(map[string]interface{})[key]
		if !exists {
			return *new(T), false
		}
	}

	if value == nil {
		return *new(T), false
	}

	v, ok := value.(T)
	if !ok {
		return *new(T), false
	}

	return v, true
}
