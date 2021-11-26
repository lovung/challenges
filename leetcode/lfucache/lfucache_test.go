package lfucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFUCache(t *testing.T) {
	// ["LFUCache","put","put","get","put","get","get","put","get","get","get"]
	// Input: [[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]
	// Expected: [null,null,null,1,null,-1,3,null,-1,3,4]
	t.Run("example 1", func(t *testing.T) {
		lfu := Constructor(2)
		lfu.Put(1, 1)
		lfu.Put(2, 2)
		assert.Equal(t, 1, lfu.Get(1))
		lfu.Put(3, 3)
		assert.Equal(t, -1, lfu.Get(2))
		assert.Equal(t, 3, lfu.Get(3))
		lfu.Put(4, 4)
		assert.Equal(t, -1, lfu.Get(1))
		assert.Equal(t, 3, lfu.Get(3))
		assert.Equal(t, 4, lfu.Get(4))
	})

	// Case 2
	// ["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]
	// Input: [[3],[2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]
	// Expected: [null,null,null,2,1,2,null,null,-1,2,1,4]
	t.Run("example 2", func(t *testing.T) {
		lfu := Constructor(3)
		lfu.Put(2, 2)
		lfu.Put(1, 1)
		assert.Equal(t, 2, lfu.Get(2))
		assert.Equal(t, 1, lfu.Get(1))
		assert.Equal(t, 2, lfu.Get(2))
		lfu.Put(3, 3)
		lfu.Put(4, 4)
		assert.Equal(t, -1, lfu.Get(3))
		assert.Equal(t, 2, lfu.Get(2))
		assert.Equal(t, 1, lfu.Get(1))
		assert.Equal(t, 4, lfu.Get(4))
	})

	// ["LFUCache","put","get"]
	// Input: [[0],[0,0],[0]]
	// Expected: [null,null,-1]
	t.Run("example 3", func(t *testing.T) {
		lfu := Constructor(0)
		lfu.Put(0, 0)
		assert.Equal(t, -1, lfu.Get(0))
	})

	// ["LFUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	// Input: [[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	// Expected: [null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]
	t.Run("example 4", func(t *testing.T) {
		lfu := Constructor(3)
		lfu.Put(1, 1)
		lfu.Put(2, 2)
		lfu.Put(3, 3)
		lfu.Put(4, 4)
		assert.Equal(t, 4, lfu.Get(4))
		assert.Equal(t, 3, lfu.Get(3))
		assert.Equal(t, 2, lfu.Get(2))
		assert.Equal(t, -1, lfu.Get(1))
		lfu.Put(5, 5)
		assert.Equal(t, -1, lfu.Get(1))
		assert.Equal(t, 2, lfu.Get(2))
		assert.Equal(t, 3, lfu.Get(3))
		assert.Equal(t, -1, lfu.Get(4))
		assert.Equal(t, 5, lfu.Get(5))
	})

	// ["LFUCache","put","put","get","put","get"]
	// Input: [[2],[1,1],[2,2],[1],[2,3],[2]]
	// Expected: [null,null,null,1,null,3]
	t.Run("example 5", func(t *testing.T) {
		lfu := Constructor(2)
		lfu.Put(1, 1)
		lfu.Put(2, 2)
		assert.Equal(t, 1, lfu.Get(1))
		lfu.Put(2, 3)
		assert.Equal(t, 3, lfu.Get(2))
	})

	// ["LFUCache","put","put","get"]
	// Input: [[1],[1,1],[2,2],[1]]
	// Expected: [null,null,null,-1]
	t.Run("example 6", func(t *testing.T) {
		lfu := Constructor(1)
		lfu.Put(1, 1)
		lfu.Put(2, 2)
		assert.Equal(t, -1, lfu.Get(1))
	})

	t.Run("panic test 1", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		var nilNode *LFUDNode
		nilNode.insertBefore(nilNode)
	})
	t.Run("panic test 2", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		var nilNode *LFUDNode
		nilNode.insertAfter(nilNode)
	})
	t.Run("panic test 3", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		var nilNode *LFUDNode
		nilNode.selfRemove()
	})
	t.Run("panic test 4", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		var nilNode *LFUDNode
		nilNode.addKey(0)
	})
	t.Run("panic test 5", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		var nilNode *LFUDNode
		nilNode.touchKey(0)
	})
	t.Run("panic test 6", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		var nilList *lfuDList
		nilList.evict()
	})
	t.Run("panic test 7", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		var nilList *lfuDList
		nilList.evict()
	})
	t.Run("panic test 8", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		var nilList *lfuDList
		nilList.addKey(0)
	})
	t.Run("panic test 9", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		nilNode := &LFUDNode{
			freq: 1,
			keys: []int{1, 2},
		}
		nilNode.touchKey(0)
	})
}
