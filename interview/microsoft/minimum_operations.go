package microsoft

import (
	llq "github.com/emirpasic/gods/v2/queues/linkedlistqueue"
)

// Change from x -> y int
// plus 1, minus 1, multi by 2

// minimum number of operation
// ---> Shortest Path

// - Can be negative

// 5 to 14
// total is 3
// - 5 + 1 = 6
// - 6 + 1 = 7
// - 7 * 2 = 14

// 5
// - 6
// - 4
// - 10
// 6
// - 5
// - 7
// - 12

type StepNode struct {
	Val   int
	Steps int
}

// BFS -> queue
func MinimumOperations(x, y int) int {
	visited := make(map[int]bool)
	q := llq.New[*StepNode]()
	if x == y {
		return 0
	}
	putQueue(x, 0, q)
	for !q.Empty() {
		n, _ := q.Dequeue()
		if visited[n.Val] {
			continue
		}
		visited[n.Val] = true
		if n.Val == y {
			return n.Steps
		}
		putQueue(n.Val, n.Steps, q)
	}
	return -1
}

// y = 14
// n         queue
// 5:0       [6:1 4:1 10:1]
// 6:1       [4:1 10:1 7:2 5:2 12:2]
// 4:1       [10:1 7:2 5:2 12:2 5:2 3:2 8:2]
// 10        [7:2 5:2 12:2 5:2 3:2 8:2]

func putQueue(x, curStep int, q *llq.Queue[*StepNode]) {
	q.Enqueue(&StepNode{
		Val:   x + 1,
		Steps: curStep + 1,
	})
	q.Enqueue(&StepNode{
		Val:   x - 1,
		Steps: curStep + 1,
	})
	q.Enqueue(&StepNode{
		Val:   x * 2,
		Steps: curStep + 1,
	})
}
