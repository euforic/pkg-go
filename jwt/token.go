package jwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	*jwt.Token
	parseError error
}

// NewToken creates a new instance of Token from a *jwt.Token
func NewToken(t *jwt.Token) *Token {
	return &Token{Token: t}
}

func (t Token) ParseError() error {
	return t.parseError
}

func (t Token) Claims() (value jwt.MapClaims, exists bool) {
	value, exists = t.Token.Claims.(jwt.MapClaims)
	return
}

func (t Token) Get(key string) (value interface{}, exists bool) {
	keys := strings.Split(key, ".")
	for i, key := range keys {
		if i == 0 {
			value, exists = t.Token.Claims.(jwt.MapClaims)[keys[0]]
			continue
		}
		value, exists = value.(map[string]interface{})[key]
	}
	return value, exists
}

// GetString returns the value associated with the key as a string.
func (t Token) GetString(key string) (s string) {
	if val, ok := t.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetBool returns the value associated with the key as a boolean.
func (t Token) GetBool(key string) (b bool) {
	if val, ok := t.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (t Token) GetInt(key string) (i int) {
	if val, ok := t.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

// GetInt64 returns the value associated with the key as an integer.
func (t Token) GetInt64(key string) (i64 int64) {
	if val, ok := t.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

// GetUint returns the value associated with the key as an unsigned integer.
func (t Token) GetUint(key string) (ui uint) {
	if val, ok := t.Get(key); ok && val != nil {
		ui, _ = val.(uint)
	}
	return
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func (t Token) GetUint64(key string) (ui64 uint64) {
	if val, ok := t.Get(key); ok && val != nil {
		ui64, _ = val.(uint64)
	}
	return
}

// GetFloat64 returns the value associated with the key as a float64.
func (t Token) GetFloat64(key string) (f64 float64) {
	if val, ok := t.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

// GetTime returns the value associated with the key as time.
func (t Token) GetTime(key string) (tm time.Time) {
	if val, ok := t.Get(key); ok && val != nil {
		tm, _ = val.(time.Time)
	}
	return
}

// GetDuration returns the value associated with the key as a duration.
func (t Token) GetDuration(key string) (d time.Duration) {
	if val, ok := t.Get(key); ok && val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

// GetSlice returns the value associated with the key as a slice of interface{}.
func (t Token) GetSlice(key string) []interface{} {
	val, ok := t.Get(key)
	if ok && val != nil {
		sv, _ := val.([]interface{})
		return sv
	}
	return []interface{}{}
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (t Token) GetStringSlice(key string) (ss []string) {
	val, ok := t.Get(key)
	if ok && val != nil {
		sv, _ := val.([]interface{})
		for _, v := range sv {
			if vv, ok := v.(string); ok {
				ss = append(ss, vv)
			}
		}
	}
	return
}
