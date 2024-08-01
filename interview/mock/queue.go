package mock

type Queue[T any] interface {
	Enqueue(T) bool
	Dequeue() (T, bool)
	Len() int
}

func New[T any](size int) Queue[T] {
	return &circleQueue[T]{
		data: make([]T, size),
	}
}

type circleQueue[T any] struct {
	data       []T
	head, tail int
	isFull     bool
}

func (q *circleQueue[T]) Len() int {
	if q.isFull {
		return len(q.data)
	}
	if q.head >= q.tail {
		return q.head - q.tail
	}
	return q.head + len(q.data) - q.tail
}

func (q *circleQueue[T]) Enqueue(v T) bool {
	if q.isFull {
		return false
	}
	q.data[q.head] = v
	q.head++
	if q.head >= len(q.data) {
		q.head = 0
	}
	if q.head == q.tail {
		q.isFull = true
	}
	return true
}

func (q *circleQueue[T]) Dequeue() (v T, ok bool) {
	if q.Len() == 0 {
		return v, false
	}
	v = q.data[q.tail]
	q.tail++
	if q.tail >= len(q.data) {
		q.tail = 0
	}
	var z T
	q.data[q.tail] = z
	if q.isFull {
		q.isFull = false
	}
	return v, true
}
