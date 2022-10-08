package grind75

// Link: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
var number2Char = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	for i := range digits {
		newChars := number2Char[digits[i]]
		if len(res) == 0 {
			for j := range newChars {
				res = append(res, string(newChars[j]))
			}
		} else {
			newRes := make([]string, 0, len(res)*len(newChars))
			for k := range res {
				for j := range newChars {
					newRes = append(newRes, res[k]+string(newChars[j]))
				}
			}
			res = newRes
		}
	}
	return res
}
