package apr2024

import (
	"github.com/lovung/ds/slice"
)

// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/
func countStudents(students []int, sandwiches []int) int {
	const (
		circular = iota
		square
	)
	n := len(students)
	circularStudents := slice.Count(students, 0)
	squareStudents := n - circularStudents
	for len(sandwiches) > 0 {
		if sandwiches[0] == students[0] {
			if sandwiches[0] == circular {
				circularStudents--
			} else {
				squareStudents--
			}
			students = slice.RemoveHead(students, 1)
			sandwiches = slice.RemoveHead(sandwiches, 1)
			continue
		}
		if sandwiches[0] == square &&
			circularStudents == len(sandwiches) {
			break
		}
		if sandwiches[0] == circular &&
			squareStudents == len(sandwiches) {
			break
		}
		slice.RotateLeft(students, 1)
	}
	return len(sandwiches)
}
