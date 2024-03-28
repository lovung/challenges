package grind75

import "github.com/lovung/ds/queue"

// Link: https://leetcode.com/problems/word-ladder/
type wordQueueItem struct {
	word    string
	stepCnt int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool)
	for i := range wordList {
		wordMap[wordList[i]] = true
	}
	visited := make(map[string]bool)
	q := queue.NewSimpleQueue[*wordQueueItem]()
	q.EnQueue(&wordQueueItem{
		word:    beginWord,
		stepCnt: 0,
	})
	for q.Len() > 0 {
		w, _ := q.DeQueue()
		if w.word == endWord {
			return w.stepCnt + 1
		}
		transform1CharWithCheck(w.word, wordMap, w.stepCnt, q, visited)
	}
	return 0
}

func transform1CharWithCheck(
	word string, wordMap map[string]bool, stepCnt int,
	q queue.Queue[*wordQueueItem], visited map[string]bool,
) {
	for i := range word {
		for j := byte(0); j < 26; j++ {
			newWord := []byte(word)
			newWord[i] = a + j
			strNewWord := string(newWord)
			if strNewWord != word && wordMap[strNewWord] && !visited[strNewWord] {
				q.EnQueue(&wordQueueItem{
					word:    strNewWord,
					stepCnt: stepCnt + 1,
				})
				visited[strNewWord] = true
			}
		}
	}
}
