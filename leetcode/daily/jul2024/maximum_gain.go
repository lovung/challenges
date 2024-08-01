package jul2024

import (
	"github.com/emirpasic/gods/v2/stacks"
	"github.com/emirpasic/gods/v2/stacks/linkedliststack"
)

// https://leetcode.com/problems/maximum-score-from-removing-substrings
func maximumGain(s string, x int, y int) int {
	stk := linkedliststack.New[rune]()
	res := 0
	for _, r := range s {
		switch r {
		case 'a':
			if x > y {
				stk.Push(r)
				continue
			}
			if val, ok := stk.Peek(); ok && val == 'b' {
				res += y
				stk.Pop()
			} else {
				stk.Push(r)
			}
		case 'b':
			if x < y {
				stk.Push(r)
				continue
			}
			if val, ok := stk.Peek(); ok && val == 'a' {
				res += x
				stk.Pop()
			} else {
				stk.Push(r)
			}
		default:
			res += cleanUp(stk, x, y)
		}
	}
	res += cleanUp(stk, x, y)
	return res
}

func cleanUp(head stacks.Stack[rune], x, y int) int {
	var (
		res  int
		tail = linkedliststack.New[rune]()
	)
	for !head.Empty() {
		h, _ := head.Pop()
		if t, ok := tail.Peek(); !ok {
			tail.Push(h)
		} else {
			if h == 'a' && t == 'b' {
				res += x
				tail.Pop()
			} else if h == 'b' && t == 'a' {
				res += y
				tail.Pop()
			} else {
				tail.Push(h)
			}
		}
	}
	return res
}

func maximumGain2(s string, x int, y int) int {
	var res int
	remove := func(score int, pattern string) {
		var stk []byte
		for i := 0; i < len(s); i++ {
			stk = append(stk, s[i])
			for len(stk) >= 2 && stk[len(stk)-2] == pattern[0] && stk[len(stk)-1] == pattern[1] {
				stk = stk[:len(stk)-2]
				res += score
			}
		}
		s = string(stk)
	}
	if x > y {
		remove(x, "ab")
		remove(y, "ba")
	} else {
		remove(y, "ba")
		remove(x, "ab")
	}
	return res
}
