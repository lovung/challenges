package jun2024

// https://leetcode.com/problems/subarray-sums-divisible-by-k/description/
func subarraysDivByK(nums []int, k int) int {
	prefixMod, result := 0, 0
	modGroups := make([]int, k)
	modGroups[0] = 1

	for i := range nums {
		prefixMod = (prefixMod + nums[i]%k + k) % k
		result += modGroups[prefixMod]
		modGroups[prefixMod]++
	}
	return result
}
