package jul2024

import (
	"strings"

	"github.com/emirpasic/gods/v2/stacks"
	"github.com/emirpasic/gods/v2/stacks/linkedliststack"
)

// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/
func reverseParentheses(s string) string {
	sb := strings.Builder{}
	all := linkedliststack.New[stacks.Stack[rune]]()
	for _, r := range s {
		switch r {
		case '(':
			newStack := linkedliststack.New[rune]()
			all.Push(newStack)
		case ')':
			top, _ := all.Pop()
			for r1, ok := top.Pop(); ok; r1, ok = top.Pop() {
				peak, ok := all.Peek()
				if ok {
					peak.Push(r1)
				} else {
					sb.WriteRune(r1)
				}
			}
		default:
			peak, ok := all.Peek()
			if ok {
				peak.Push(r)
			} else {
				sb.WriteRune(r)
			}
		}
	}
	return sb.String()
}
