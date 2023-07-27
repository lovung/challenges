package grind75

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{
		data: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	e := this.data[0]
	this.data = this.data[1:]
	return e
}

func (this *MyQueue) Peek() int {
	return this.data[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}
