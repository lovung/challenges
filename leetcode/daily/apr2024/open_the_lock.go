package apr2024

import (
	"strconv"

	"github.com/lovung/ds/queue"
)

// https://leetcode.com/problems/open-the-lock
func openLock(deadends []string, target string) int {
	targetNum, _ := strconv.Atoi(target)
	visited := make([]bool, 1e5)
	for i := range deadends {
		s, _ := strconv.Atoi(deadends[i])
		visited[s] = true
	}
	if visited[0] {
		return -1
	}
	visited[0] = true
	type step struct {
		s   int
		cnt int
	}
	const src = 0
	q := queue.NewCircularQueue[step](1e5)
	q.EnQueue(step{src, 0})
	for q.Len() > 0 {
		cur, _ := q.DeQueue()
		if cur.s == targetNum {
			return cur.cnt
		}
		nexts := genNextTurn(cur.s)
		for i := range nexts {
			if !visited[nexts[i]] {
				q.EnQueue(step{nexts[i], cur.cnt + 1})
				visited[nexts[i]] = true
			}
		}
	}
	return -1
}

var ten = []int{1, 10, 100, 1000, 10000}

func genNextTurn(cur int) []int {
	res := make([]int, 0, 8)
	for i := 0; i < 4; i++ {
		b := (cur % ten[i+1]) / ten[i]
		res = append(res, cur-b*ten[i]+increase(b)*ten[i])
		res = append(res, cur-b*ten[i]+decrease(b)*ten[i])
	}
	return res
}

func increase(b int) int {
	if b == 9 {
		return 0
	}
	return b + 1
}

func decrease(b int) int {
	if b == 0 {
		return 9
	}
	return b - 1
}
