package grind75

import (
	"strconv"

	"github.com/lovung/ds/stack"
)

// Link: https://leetcode.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stck := stack.NewArrayStack[int](1e4)
	for i := range tokens {
		switch tokens[i] {
		case "+":
			num1 := stck.Pop()
			num2 := stck.Pop()
			stck.Push(num2 + num1)
		case "*":
			num1 := stck.Pop()
			num2 := stck.Pop()
			stck.Push(num2 * num1)
		case "-":
			num1 := stck.Pop()
			num2 := stck.Pop()
			stck.Push(num2 - num1)
		case "/":
			num1 := stck.Pop()
			num2 := stck.Pop()
			stck.Push(num2 / num1)
		default:
			num, _ := strconv.Atoi(tokens[i])
			stck.Push(num)
		}
	}
	return stck.Pop()
}
