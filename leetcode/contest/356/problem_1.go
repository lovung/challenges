package contest356

// https://leetcode.com/problems/number-of-employees-who-met-the-target/
func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	cnt := 0
	for i := range hours {
		if hours[i] >= target {
			cnt++
		}
	}
	return cnt
}
