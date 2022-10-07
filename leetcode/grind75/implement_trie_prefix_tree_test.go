package grind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Link: https://leetcode.com/problems/implement-trie-prefix-tree/
func TestTrie(t *testing.T) {
	t.Run("apple", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("apple")
		assert.True(t, trie.Search("apple"))
		assert.False(t, trie.Search("app"))
		assert.True(t, trie.StartsWith("app"))
		trie.Insert("app")
		assert.True(t, trie.Search("app"))
		assert.False(t, trie.Search("applepay"))
	})
	t.Run("a", func(t *testing.T) {
		trie := NewTrie()
		assert.False(t, trie.StartsWith("a"))
	})
}
