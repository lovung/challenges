package microsoft

import "sort"

// Link: https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique
func minDeletions(s string) int {
	const a = 'a'
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-a]++
	}
	ret := 0
	sort.Ints(freq)

	for i := 25; freq[i] > 0 && i > 0; i-- {
		j := i - 1
		if freq[i] != freq[j] {
			continue
		}
		for ; j >= 0 && freq[i] == freq[j]; j-- {
			ret++
			freq[j]--
		}
	}
	return ret
}
