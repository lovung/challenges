package may2024

import (
	"slices"
)

const a = 'a'

// https://leetcode.com/problems/maximum-score-words-formed-by-letters/
func maxScoreWords(words []string, letters []byte, score []int) int {
	letterCnt := letterCount(letters)
	pairs := calcWordScore(words, letterCnt, score)
	slices.SortFunc(pairs, func(a, b *pair) int { return b.score - a.score })
	totalScore := 0
	for i := range pairs {
		totalScore += pairs[i].score
	}
	res := 0
	recursiveMaxScoreWords(pairs, 0, 0, totalScore, &res, letterCnt)
	return res
}

func recursiveMaxScoreWords(pairs []*pair, idx, curScore, remainScore int, res *int, cnt []int) {
	if idx >= len(pairs) || pairs[idx].score == 0 {
		return
	}
	// Early return
	if curScore+remainScore < *res {
		return
	}
	takeCnt := slices.Clone(cnt)
	if remove(pairs[idx].word, takeCnt) {
		*res = max(*res, curScore+pairs[idx].score)
		recursiveMaxScoreWords(pairs, idx+1, curScore+pairs[idx].score, remainScore-pairs[idx].score, res, takeCnt)
	}
	recursiveMaxScoreWords(pairs, idx+1, curScore, remainScore-pairs[idx].score, res, cnt)
}

type pair struct {
	word  string
	score int
}

func calcWordScore(words []string, cnt []int, score []int) []*pair {
	pairs := make([]*pair, len(words))
	for i := range words {
		curLettterCnt := letterCount([]byte(words[i]))
		if !canFit(curLettterCnt, cnt) {
			pairs[i] = &pair{words[i], 0}
		} else {
			pairs[i] = &pair{words[i], countScore(words[i], score)}
		}
	}
	return pairs
}

func remove(s string, cnt []int) bool {
	for i := range s {
		if cnt[s[i]-a] == 0 {
			return false
		}
		cnt[s[i]-a]--
	}
	return true
}

func letterCount(letters []byte) []int {
	cnt := [26]int{}
	for i := range letters {
		cnt[letters[i]-a]++
	}
	return cnt[:]
}

func countScore(s string, score []int) int {
	res := 0
	for i := range s {
		res += score[s[i]-a]
	}
	return res
}

func canFit(need, limit []int) bool {
	for i := range 26 {
		if need[i] > limit[i] {
			return false
		}
	}
	return true
}
