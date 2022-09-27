package grind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cloneGraph(t *testing.T) {
	t.Run("square", func(t *testing.T) {
		n1 := &Node{Val: 1}
		n2 := &Node{Val: 2}
		n3 := &Node{Val: 3}
		n4 := &Node{Val: 4}
		n1.Neighbors = []*Node{n2, n4}
		n2.Neighbors = []*Node{n1, n3}
		n3.Neighbors = []*Node{n2, n4}
		n4.Neighbors = []*Node{n1, n3}

		newNode := cloneGraph(n1)
		assert.True(t, compareNodes(n1, newNode))
		assert.False(t, compareNodes(n1, n1))
	})
	t.Run("nil", func(t *testing.T) {
		cloneGraph(nil)
	})
}
