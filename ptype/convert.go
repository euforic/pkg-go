package ptype

func Deref[T any](in *T) T {
	if in == nil {
		out := new(T)
		return *out
	}

	return *in
}

func Ptr[T any](in T) *T {
	return &in
}
