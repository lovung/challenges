package grind75

//              2    3    3
// take   0/0   2/2  3    max(0+3, 2)
// dont   0/0   0/0  0/2  max(2, 3)
// -> 3

//          2    1    5               3
// take   0 2/2  1/1  5/7             3+1
// dont   0 0/0  0/2  1/2             7

type can struct {
	min, max int
}

// https://leetcode.com/problems/house-robber-ii/
func rob2_1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	take, dont := can{nums[1], nums[1]}, can{0, nums[0]}
	for i := 2; i < n-1; i++ {
		take, dont = dont.add(nums[i]), dont.merge(take)
	}
	lastTake := max(dont.min+nums[n-1], dont.max)
	lastDont := maxOf(take, dont)
	return max(lastTake, lastDont)
}

func (c can) add(n int) can {
	return can{c.min + n, c.max + n}
}

func (c can) merge(n can) can {
	return can{max(c.min, n.min), max(c.max, n.max)}
}

func maxOf(case1, case2 can) int {
	return max(case1.max, case2.max)
}

func rob2_2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a := rob(nums[:len(nums)-1])
	b := rob(nums[1:])
	return max(a, b)
}
