package contest

const a = 'a'
const n = 26

func checkDistances(s string, distance []int) bool {
	lastIndex := make([]int, 0, n)
	for i := 0; i < n; i++ {
		lastIndex = append(lastIndex, -1)
	}
	for i := range s {
		if lastIndex[s[i]-a] >= 0 {
			if i-lastIndex[s[i]-a] != distance[s[i]-a]+1 {
				return false
			}
		} else {
			lastIndex[s[i]-a] = i
		}
	}
	return true
}
