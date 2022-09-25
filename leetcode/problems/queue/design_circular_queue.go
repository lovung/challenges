package queue

// Link: https://leetcode.com/problems/design-circular-queue/
type MyCircularQueue struct {
	buffer []int
	size   int

	head, tail int
	len        int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		buffer: make([]int, k),
		size:   k,
		head:   0,
		tail:   -1,
		len:    0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.tail++
	if this.tail == this.size {
		this.tail = 0
	}
	this.buffer[this.tail] = value
	this.len++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head++
	if this.head == this.size {
		this.head = 0
	}
	this.len--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.len == this.size
}
