package oct2022

import (
	"testing"

	"github.com/lovung/ds/lists"
	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		head := lists.NewLinkedListFromValues([]int{4, 5, 1, 9})
		deleteNode(head.Next)
		assert.Equal(t, lists.NewLinkedListFromValues([]int{4, 1, 9}), head)
	})
	t.Run("next-next", func(t *testing.T) {
		head := lists.NewLinkedListFromValues([]int{4, 5, 1, 9})
		deleteNode(head.Next.Next)
		assert.Equal(t, lists.NewLinkedListFromValues([]int{4, 5, 9}), head)
	})
}
