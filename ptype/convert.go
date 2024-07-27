package ptype

// Deref returns the value that in points to, or a zero value if in is nil.
func Deref[T any](in *T) T { //nolint:ireturn
	if in == nil {
		out := new(T)

		return *out
	}

	return *in
}

// Ptr returns a pointer to in.
func Ptr[T any](in T) *T {
	return &in
}
