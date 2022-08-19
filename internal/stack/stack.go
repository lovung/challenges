package stack

type Stack[T any] interface {
	Push(v T)
	Pop() T
	Peek() T
	Clear()
	Len() int
	All() []T
}
