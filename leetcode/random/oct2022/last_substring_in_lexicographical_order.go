package oct2022

// Link: https://leetcode.com/problems/last-substring-in-lexicographical-order/
func lastSubstring(s string) string {
	// First idea:
	// - Find the greatest (by byte) char in the string. 						: O(N)
	// - If only 1 index, return longest substring (from this index to the end)
	// - If many indexs (like: "theleetcode", 't' at 0 and 6)
	// 		-> put all substring from these index to the list 					: O(N)
	//			* no need to check the shorter substrings because
	//				it will be lexicographically smaller the longest
	//		->  found the max substring	by check all the list					: O(N)
	var countIndex [26][]int
	const a = 'a'
	const z = 'a'

	// - Find the greatest (by byte) char in the string.
	for i := range s {
		countIndex[s[i]-a] = append(countIndex[s[i]-a], i)
	}

	greatestChar := 25
	for ; greatestChar >= 0; greatestChar-- {
		if len(countIndex[greatestChar]) > 0 {
			break
		}
	}

	// - If only 1 index, return longest substring (from this index to the end)
	if len(countIndex[greatestChar]) == 1 {
		return s[countIndex[greatestChar][0]:]
	}

	// - If many indexs
	// ->  found the max substring	by check all the list
	maxSubString := ""
	for _, index := range countIndex[greatestChar] {
		if s[index:] > maxSubString {
			maxSubString = s[index:]
		}
	}

	return maxSubString
}

func lastSubstring2(s string) string {
	i, j, k := 0, 1, 0
	n := len(s)
	for j+k < n {
		switch {
		case s[i+k] == s[j+k]:
			k++
			continue
		case s[i+k] > s[j+k]:
			j += k + 1
		default:
			i += k + 1
		}
		if i == j {
			j++
		}
		k = 0
	}
	return s[i:]
}
