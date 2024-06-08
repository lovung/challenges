package biweekly131

// https://leetcode.com/problems/find-the-xor-of-numbers-which-appear-twice/submissions/
func queryResults(_ int, queries [][]int) []int {
	type ball int
	type color int
	res := make([]int, len(queries))
	colorMap := make(map[ball]color)
	colorCnt := make(map[color]int)
	distinct := 0
	for i, q := range queries {
		// new color
		if oldColor, ok := colorMap[ball(q[0])]; ok {
			colorCnt[oldColor]--
			if colorCnt[oldColor] == 0 {
				distinct--
			}
		}
		colorMap[ball(q[0])] = color(q[1])
		colorCnt[color(q[1])]++
		if colorCnt[color(q[1])] == 1 {
			distinct++
		}
		res[i] = distinct
	}
	return res
}
