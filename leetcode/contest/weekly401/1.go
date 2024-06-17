package weekly401

// https://leetcode.com/problems/find-the-child-who-has-the-ball-after-k-seconds/description/
func numberOfChild(n int, k int) int {
	cur, dir := 0, true
	for range k {
		if dir {
			cur++
			if cur == n {
				cur -= 2
				dir = !dir
			}
		} else {
			cur--
			if cur < 0 {
				cur += 2
				dir = !dir
			}
		}
	}
	return cur
}
