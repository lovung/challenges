package medium

// Link: https://leetcode.com/problems/interval-list-intersections/
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0)
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		if firstList[i][1] < secondList[j][0] {
			i++
			continue
		}
		if firstList[i][0] > secondList[j][1] {
			j++
			continue
		}
		res = append(res, []int{
			max(firstList[i][0], secondList[j][0]),
			min(firstList[i][1], secondList[j][1]),
		})
		switch {
		case firstList[i][1] > secondList[j][1]:
			j++
		case firstList[i][1] < secondList[j][1]:
			i++
		default:
			i++
			j++
		}
	}
	return res
}
