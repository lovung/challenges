package grind75

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{
		data: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.data = append(q.data, x)
}

func (q *MyQueue) Pop() int {
	e := q.data[0]
	q.data = q.data[1:]
	return e
}

func (q *MyQueue) Peek() int {
	return q.data[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.data) == 0
}
