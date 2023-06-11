package mapsi

import (
	"errors"
	"strings"
)

// Value is an interface that can be used to store any value in a map.
type Result interface{}

// Get gets the value of a key in a map. If the key does not exist, it returns an error.
func Get(m map[string]interface{}, key string) (Result, error) {
	// Split the key into head and the rest
	segments := strings.SplitN(key, ".", 2)
	head := segments[0]

	// Get value associated with head
	value, ok := m[head]
	if !ok {
		return nil, errors.New("key does not exist")
	}

	// If key has no dot, return value
	if len(segments) == 1 {
		return value, nil
	}

	// If value is a nested map, call Get recursively
	if nestedMap, ok := value.(map[string]interface{}); ok {
		return Get(nestedMap, segments[1])
	}

	// If value is not a nested map and key has dot, return error
	return nil, errors.New("key does not exist")
}

// Set sets the value of a key in a map. If the key does not exist, it is created.
func Set(m map[string]interface{}, key string, value interface{}) error {
	segments := strings.SplitN(key, ".", 2)
	head := segments[0]

	if len(segments) == 1 {
		m[head] = value
		return nil
	}

	if nestedMap, ok := m[head].(map[string]interface{}); ok {
		return Set(nestedMap, segments[1], value)
	}

	// if head does not exist, create it
	m[head] = make(map[string]interface{})
	if nestedMap, ok := m[head].(map[string]interface{}); ok {
		return Set(nestedMap, segments[1], value)
	}

	return errors.New("key does not exist")
}
