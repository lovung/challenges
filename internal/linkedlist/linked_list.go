package linkedlist

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func NewListNode[T any](value T) ListNode[T] {
	return ListNode[T]{
		Val: value,
	}
}

func NewLinkedListFromValues[T any](values []T) *ListNode[T] {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode[T]{
		Val: values[0],
	}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode[T]{
			Val: values[i],
		}
		current = current.Next
	}
	return head
}
