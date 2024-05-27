package may2024

import (
	"github.com/lovung/ds/queue"
	"github.com/lovung/ds/tries"
)

// https://leetcode.com/problems/word-break-ii/
func wordBreak(s string, wordDict []string) []string {
	trie := tries.NewSimpleTrie()
	for i := range wordDict {
		trie.Insert(wordDict[i])
	}
	res := make([]string, 0, 1<<5)
	type pair struct {
		idx int
		cur string
	}
	q := queue.NewSimpleQueue[pair]()
	q.EnQueue(pair{0, ""})
	for q.Len() > 0 {
		item, _ := q.DeQueue()
		l := item.idx
		for r := l + 1; r <= len(s); r++ {
			if !trie.StartsWith(s[l:r]) {
				break
			}
			if trie.Search(s[l:r]) {
				newCur := s[l:r]
				if len(item.cur) > 0 {
					newCur = item.cur + " " + s[l:r]
				}
				if r == len(s) {
					res = append(res, newCur)
				} else {
					q.EnQueue(pair{r, newCur})
				}
			}
		}
	}
	return res
}
