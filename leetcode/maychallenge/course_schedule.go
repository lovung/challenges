package maychallenge

// Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/538/week-5-may-29th-may-31st/3344/
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		graph = make([][]int, numCourses)
		dfs   func(int, []int) bool
	)
	for _, e := range prerequisites {
		graph[e[0]] = append(graph[e[0]], e[1])
	}
	dfs = func(n int, ref []int) bool {
		ok := true
		for i := 0; (i < len(graph[n])) && ok; i++ {
			for _, e := range ref {
				if graph[n][i] == e {
					return false
				}
			}
			ref = append(ref, n)
			ok = dfs(graph[n][i], ref)
		}
		return ok
	}

	for i := 0; i < numCourses; i++ {
		var ref []int
		if !dfs(i, ref) {
			return false
		}
	}
	return true
}
