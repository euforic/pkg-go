package sliceutil

import (
	"sort"
)

// Compare will check if two slices are equal
// even if they aren't in the same order
func Compare[K comparable](s1 []K, s2 []K) bool {
	if s1 == nil || s2 == nil {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	// setup maps to store values and count of slices
	m1 := make(map[K]int)
	m2 := make(map[K]int)

	for i := 0; i < len(s1); i++ {
		// Add each value to map and increment for each found
		m1[s1[i]]++
		m2[s2[i]]++
	}

	for key := range m1 {
		if m1[key] != m2[key] {
			return false
		}
	}

	return true
}

// OrderedCompare will check if two slices are equal, taking order into consideration.
func OrderedCompare[K comparable](s1 []K, s2 []K) bool {
	//If both are nil, they are equal
	if s1 == nil && s2 == nil {
		return true
	}

	//If only one is nil, they are not equal (!= represents XOR)
	if (s1 == nil) != (s2 == nil) {
		return false
	}

	//If both are nil, they are equal
	if s1 == nil || s2 == nil {
		return false
	}

	//If the lengths are different, the slices are not equal
	if len(s1) != len(s2) {
		return false
	}

	//Loop through and compare the slices at each index
	for i := 0; i < len(s1); i++ {
		if s2[i] != s1[i] {
			return false
		}
	}

	//If nothing has failed up to this point, the slices are equal
	return true
}

// Contains checks if a slice contains an element
func Contains[K comparable](s []K, e K) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Reverse reverses slices of any type
func Reverse[K comparable](s []K) []K {
	a := make([]K, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

// FastContains offers a faster implementation of contains
func FastContains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}
