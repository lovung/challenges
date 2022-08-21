package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
func TestSmallestInfiniteSet(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		// ["SmallestInfiniteSet","addBack","popSmallest","popSmallest","popSmallest","addBack","popSmallest","popSmallest","popSmallest"]
		// [[],[2],[],[],[],[1],[],[],[]]
		obj := Constructor()
		assert.Equal(t, 1, obj.PopSmallest())
		obj.AddBack(2)
		assert.Equal(t, 2, obj.PopSmallest())
		assert.Equal(t, 3, obj.PopSmallest())
		assert.Equal(t, 4, obj.PopSmallest())
		obj.AddBack(1)
		assert.Equal(t, 1, obj.PopSmallest())
		assert.Equal(t, 5, obj.PopSmallest())
	})
}
