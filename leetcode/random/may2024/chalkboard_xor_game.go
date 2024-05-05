package may2024

// https://leetcode.com/problems/chalkboard-xor-game/
func xorGame(nums []int) bool {
	xor := 0
	for i := range nums {
		xor ^= nums[i]
	}
	if xor == 0 {
		return true
	}
	return len(nums)%2 == 0
}
