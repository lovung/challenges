package may2024

// https://leetcode.com/problems/score-after-flipping-matrix/
func matrixScore2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	nums := nums(grid)

	for i := range nums {
		if nums[i]&(1<<(n-1)) == 0 {
			flipNum(&nums[i], n)
		}
	}
	for j := range n {
		if countZeroBits(nums, j) > m/2 {
			flipColNums(nums, j)
		}
	}
	res := 0
	for i := range nums {
		res += nums[i]
	}
	return res
}

func nums(grid [][]int) []int {
	res := make([]int, 0, len(grid))
	for i := range grid {
		num, bit := 0, 1
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				num |= bit
			}
			bit <<= 1
		}
		res = append(res, num)
	}
	return res
}

func flipNum(num *int, n int) {
	bit := 1
	for c := n - 1; c >= 0; c-- {
		*num ^= bit
		bit <<= 1
	}
}

func flipColNums(nums []int, b int) {
	bit := 1
	for range b {
		bit <<= 1
	}
	for r := range nums {
		nums[r] ^= bit
	}
}
func countZeroBits(nums []int, b int) int {
	bit, cnt := 1, 0
	for range b {
		bit <<= 1
	}
	for r := range nums {
		if nums[r]&bit == 0 {
			cnt++
		}
	}
	return cnt
}
