package hard

// Link: https://leetcode.com/problems/total-appeal-of-a-string/
const (
	a    = 'a'
	atoz = 26
)

// Time: O(n)
// Space: O(1)
func appealSum2(s string) int64 {
	sum := 0
	colSum := 0
	latestIndex := [atoz]int{}
	for i := range latestIndex {
		latestIndex[i] = -1
	}
	for i := range s {
		// passed len - last index
		colSum += i - latestIndex[s[i]-a]
		sum += colSum
		latestIndex[s[i]-a] = i
	}
	return int64(sum)
}

// Time: O(n^2)
// Space: O(n)
func appealSum(s string) int64 {
	l := len(s)
	sum := l
	currentMark := [atoz]int{}
	for i := range s {
		currentMark[s[i]-a]++
		// skip first charactor because we have process as assigning 1 above
		if i == 0 {
			continue
		}
		newMark := [atoz]int{}
		copy(newMark[:], currentMark[:])
		sum += countAppeal(newMark[:])
		for j := 1; j < l-i; j++ {
			newMark[s[j-1]-a]--
			newMark[s[i+j]-a]++
			sum += countAppeal(newMark[:])
		}
	}
	return int64(sum)
}

// O(26) ~ O(1)
func countAppeal(mark []int) int {
	count := 0
	for i := range mark {
		if mark[i] > 0 {
			count++
		}
	}
	return count
}
