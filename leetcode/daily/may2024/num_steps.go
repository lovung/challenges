package may2024

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/
func numSteps(org string) int {
	res := 0
	s := []byte(org)
	for !(len(s) == 1 && s[0] == '1') {
		res++
		if s[len(s)-1] == '0' { // even number
			s = s[:len(s)-1]
			continue
		}
		// odd number need to be add 1
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '0' {
				s[i]++
				break
			}
			// 1 + 1 = 10 -> mean convert it to '0' and process next digit
			s[i] = '0'
			if i == 0 {
				// if biggest digit, add 1 more char
				s = append([]byte{'1'}, s...)
			}
		}
	}
	return res
}
