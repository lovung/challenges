package grind75

import (
	"github.com/lovung/ds/queue"
	"github.com/lovung/ds/tries"
)

// Link: https://leetcode.com/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	wordDictTrie := tries.NewSimpleTrie()
	for i := range wordDict {
		wordDictTrie.Insert(wordDict[i])
	}
	q := queue.NewSimpleQueue[int]()
	q.EnQueue(0)
	processedIndex := make([]bool, len(s))
	processedIndex[0] = true
	for q.Len() > 0 {
		start, _ := q.DeQueue()
		for end := start + 1; end <= len(s); end++ {
			if !wordDictTrie.StartsWith(s[start:end]) {
				break
			}
			if wordDictTrie.Search(s[start:end]) {
				if end == len(s) {
					return true
				} else if !processedIndex[end] {
					q.EnQueue(end)
					processedIndex[end] = true
				}
			}
		}
	}
	return false
}
