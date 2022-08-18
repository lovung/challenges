package pointer

func N[T any](value T) *T {
	return &value
}
