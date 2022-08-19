package stack

type ArrayStack[T any] struct {
	data []T
	cap  int
}

func NewArrayStack[T any](cap int) Stack[T] {
	return &ArrayStack[T]{
		data: make([]T, 0, cap),
		cap:  cap,
	}
}

func (s *ArrayStack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *ArrayStack[T]) Pop() (v T) {
	if len(s.data) == 0 {
		return v
	}
	idx := len(s.data) - 1
	v = s.data[idx]
	s.data = s.data[:idx]
	return v
}

func (s *ArrayStack[T]) Clear() {
	s.data = make([]T, 0, s.cap)
}

func (s *ArrayStack[T]) Peek() (v T) {
	if len(s.data) == 0 {
		return v
	}
	idx := len(s.data) - 1
	return s.data[idx]
}

func (s *ArrayStack[T]) Len() int {
	return len(s.data)
}

func (s *ArrayStack[T]) All() []T {
	return s.data
}
