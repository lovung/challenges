package problems

/*
 * Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3318/
 */

func canConstruct(ransomNote string, magazine string) bool {
	lenRansomNote := len(ransomNote)
	lenMagazine := len(magazine)
	if lenRansomNote > lenMagazine {
		return false
	}
	array := make([]int, 26)
	for _, e := range magazine {
		array[e-'a']++
	}
	for _, e := range ransomNote {
		if array[e-'a'] > 0 {
			array[e-'a']--
		} else {
			return false
		}
	}
	return true
}
