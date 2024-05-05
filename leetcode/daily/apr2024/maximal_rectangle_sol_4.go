package apr2024

import "github.com/lovung/ds/stack"

// https://leetcode.com/problems/maximal-rectangle/
func maximalRectangle4(matrix [][]byte) int {
	m := len(matrix[0])
	v := make([]int, m)
	lf := make([]int, m)
	rg := make([]int, m)
	s := stack.NewArrayStack[int](m)
	ans := 0
	for _, row := range matrix {
		for i, c := range row {
			v[i] = If(c == '1', v[i]+1, 0)
		}

		for i := range row {
			for s.Len() != 0 && v[s.Peek()] >= v[i] {
				s.Pop()
			}
			lf[i] = If(s.Len() == 0, -1, s.Peek())
			s.Push(i)
		}
		s.Clear()

		for i := m - 1; i >= 0; i-- {
			for s.Len() != 0 && v[s.Peek()] >= v[i] {
				s.Pop()
			}
			rg[i] = If(s.Len() == 0, m, s.Peek())
			s.Push(i)
		}
		s.Clear()

		for i := range m {
			l := lf[i] + 1
			r := rg[i] - 1
			area := (r - l + 1) * v[i]
			ans = max(ans, area)
		}
	}
	return ans
}

func If[T any](cond bool, trueVal, falseVal T) T {
	if cond {
		return trueVal
	}
	return falseVal
}
