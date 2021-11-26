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

	// ["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
	// Input: [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
	// Expected: [null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]
	t.Run("example 7", func(t *testing.T) {
		lfu := Constructor(10)
		lfu.Put(10, 13)
		lfu.Put(3, 17)
		lfu.Put(6, 11)
		lfu.Put(10, 5)
		lfu.Put(9, 10)
		assert.Equal(t, -1, lfu.Get(13))
		lfu.Put(2, 19)
		assert.Equal(t, 19, lfu.Get(2))
		assert.Equal(t, 17, lfu.Get(3))
		lfu.Put(5, 25)
		assert.Equal(t, -1, lfu.Get(8))
		lfu.Put(9, 22)
		lfu.Put(5, 5)
		lfu.Put(1, 30)
		assert.Equal(t, -1, lfu.Get(11))
		lfu.Put(9, 12)
		assert.Equal(t, -1, lfu.Get(7))
		assert.Equal(t, 5, lfu.Get(5))
		assert.Equal(t, -1, lfu.Get(8))
		assert.Equal(t, 12, lfu.Get(9))
		lfu.Put(4, 30)
		lfu.Put(9, 3)
		assert.Equal(t, 3, lfu.Get(9))
		assert.Equal(t, 5, lfu.Get(10))
		assert.Equal(t, 5, lfu.Get(10))
		lfu.Put(6, 14)
		lfu.Put(3, 1)
		assert.Equal(t, 1, lfu.Get(3))
		lfu.Put(10, 11)
		assert.Equal(t, -1, lfu.Get(8))
		lfu.Put(2, 14)
		assert.Equal(t, 30, lfu.Get(1))
		assert.Equal(t, 5, lfu.Get(5))
		assert.Equal(t, 30, lfu.Get(4))
		lfu.Put(11, 4)
		lfu.Put(12, 24)
		lfu.Put(5, 18)
		assert.Equal(t, -1, lfu.Get(13))
		lfu.Put(7, 23)
		assert.Equal(t, -1, lfu.Get(8))
		assert.Equal(t, 24, lfu.Get(12))
		lfu.Put(3, 27)
		lfu.Put(2, 12)
		assert.Equal(t, 18, lfu.Get(5))

		lfu.Put(2, 9)
		lfu.Put(13, 4)
		lfu.Put(8, 18)
		lfu.Put(1, 7)
		assert.Equal(t, 14, lfu.Get(6))

		lfu.Put(9, 29)
		lfu.Put(8, 21)
		assert.Equal(t, 18, lfu.Get(5))
		lfu.Put(6, 30)
		lfu.Put(1, 12)
		assert.Equal(t, 11, lfu.Get(10))

		lfu.Put(4, 15)
		lfu.Put(7, 22)
		lfu.Put(11, 26)
		lfu.Put(8, 17)
		lfu.Put(9, 29)
		assert.Equal(t, 18, lfu.Get(5))

		lfu.Put(3, 4)
		lfu.Put(11, 30)
		assert.Equal(t, -1, lfu.Get(12))
		lfu.Put(4, 29)
		assert.Equal(t, 4, lfu.Get(3))
	})

	t.Run("example 7", func(t *testing.T) {
		lfu := Constructor(10)
		lfu.Put(1, 1)
		lfu.Put(2, 2)
		lfu.Put(3, 3)
		assert.Equal(t, 1, lfu.Get(1))
		lfu.Put(3, 4)
		assert.Equal(t, 4, lfu.Get(3))
	})
}
