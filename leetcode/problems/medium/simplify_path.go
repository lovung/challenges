package medium

import (
	"strings"

	"github.com/lovung/challenges/internal/stack"
)

const sep = "/"

// Link: https://leetcode.com/problems/simplify-path/
func simplifyPath(path string) string {
	pathStack := stack.NewArrayStack[string](100)
	paths := strings.Split(path, sep)
	for _, p := range paths {
		switch p {
		case ".", "":
			continue
		case "..":
			pathStack.Pop()
		default:
			pathStack.Push(p)
		}
	}
	return sep + strings.Join(pathStack.All(), sep)
}
