package grind75

import (
	"testing"

	"github.com/lovung/ds/lists"
	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		n1 := lists.NewListNode(3)
		n2 := lists.NewListNode(2)
		n3 := lists.NewListNode(0)
		n4 := lists.NewListNode(4)

		n1.Next = &n2
		n2.Next = &n3
		n3.Next = &n4
		n4.Next = &n2

		assert.True(t, hasCycle(&n1))
	})
	t.Run("2", func(t *testing.T) {
		n1 := lists.NewListNode(1)
		n2 := lists.NewListNode(2)

		n1.Next = &n2
		n2.Next = &n1

		assert.True(t, hasCycle(&n1))
	})

	t.Run("3", func(t *testing.T) {
		n1 := lists.NewListNode(1)

		assert.False(t, hasCycle(&n1))
	})

	t.Run("4", func(t *testing.T) {
		n1 := lists.NewListNode(1)
		n1.Next = &n1

		assert.True(t, hasCycle(&n1))
	})
}
