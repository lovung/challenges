package stack

// Stack is the FILO container
// helps holding values
type Stack[T any] interface {
	Push(v T)
	Pop() T
	Peek() T
	Clear()
	Len() int
	All() []T
}
