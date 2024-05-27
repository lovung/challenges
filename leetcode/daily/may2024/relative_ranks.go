package may2024

import (
	"sort"
	"strconv"
)

// https://leetcode.com/problems/relative-ranks/
func findRelativeRanks(score []int) []string {
	type ranking struct {
		score int
		index int
		rank  int
	}

	rankings := make([]ranking, len(score))
	for i, s := range score {
		rankings[i] = ranking{s, i, 0}
	}
	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].score > rankings[j].score
	})
	for i := range rankings {
		rankings[i].rank = i + 1
	}
	res := make([]string, len(score))
	for i := range rankings {
		switch rankings[i].rank {
		case 1:
			res[rankings[i].index] = "Gold Medal"
		case 2:
			res[rankings[i].index] = "Silver Medal"
		case 3:
			res[rankings[i].index] = "Bronze Medal"
		default:
			res[rankings[i].index] = strconv.Itoa(rankings[i].rank)
		}
	}
	return res
}
