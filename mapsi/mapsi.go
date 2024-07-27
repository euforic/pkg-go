package mapsi

import (
	"errors"
	"strings"
)

var (
	// ErrKeyDoesNotExist is returned when a key does not exist in a map.
	ErrKeyDoesNotExist = errors.New("key does not exist")
	// ErrAssertionFailed is returned when a type assertion fails.
	ErrAssertionFailed = errors.New("type assertion failed")
)

// Get gets the value of a key in a map. If the key does not exist, it returns an error.
func Get[T any](m map[string]any, key string) (T, error) { //nolint:ireturn
	var zero T
	segments := strings.SplitN(key, ".", 2) //nolint:mnd
	head := segments[0]

	// Get value associated with head
	value, ok := m[head]
	if !ok {
		return zero, ErrKeyDoesNotExist
	}

	// If key has no dot, return value
	if len(segments) == 1 {
		if v, ok := value.(T); ok {
			return v, nil
		}

		return zero, ErrAssertionFailed
	}

	// If value is a nested map, call Get recursively
	if nestedMap, ok := value.(map[string]any); ok {
		return Get[T](nestedMap, segments[1])
	}

	// If value is not a nested map and key has dot, return error
	return zero, ErrKeyDoesNotExist
}

// Set sets the value of a key in a map. If the key does not exist, it is created.
func Set[T any](m map[string]any, key string, value T) error {
	segments := strings.SplitN(key, ".", 2) //nolint:mnd
	head := segments[0]

	if len(segments) == 1 {
		m[head] = value

		return nil
	}

	if nestedMap, ok := m[head].(map[string]any); ok {
		return Set(nestedMap, segments[1], value)
	}

	// if head does not exist, create it
	m[head] = make(map[string]any)
	if nestedMap, ok := m[head].(map[string]any); ok {
		return Set(nestedMap, segments[1], value)
	}

	return ErrKeyDoesNotExist
}
