package jul2024

// https://leetcode.com/problems/pass-the-pillow/
func passThePillow(n int, time int) int {
	step := n - 1
	if (time/step)%2 != 0 {
		return n - time%step
	}
	return 1 + time%step
}
