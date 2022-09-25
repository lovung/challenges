package grind75

// Link: https://leetcode.com/problems/course-schedule/
func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
	}
	checkOKCourses := make([]bool, numCourses)
	checkingCircle := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !dfsCheckCircle(i, adj, checkOKCourses, checkingCircle) {
			return false
		}
	}
	return true
}

func dfsCheckCircle(course int, adj [][]int, checkOKCourses, checkingCircle []bool) bool {
	if checkOKCourses[course] {
		return true
	}
	if checkingCircle[course] == true {
		return false
	}
	checkingCircle[course] = true
	preCheck := true
	for _, i := range adj[course] {
		preCheck = preCheck && dfsCheckCircle(i, adj, checkOKCourses, checkingCircle)
	}
	if preCheck {
		checkOKCourses[course] = true
	}
	return preCheck
}
