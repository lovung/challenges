package problems

func possibleBipartition(N int, dislikes [][]int) bool {
	var (
		graph = make([][]int, N+1)
		color = make(map[int]int)
		dfs   func(int, int) bool
	)
	for _, e := range dislikes {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	dfs = func(node, c int) bool {
		if val, ok := color[node]; ok {
			return val == c
		}
		color[node] = c
		for _, e := range graph[node] {
			if !dfs(e, c^1) {
				return false
			}
		}
		return true
	}

	for i := 1; i <= N; i++ {
		if _, ok := color[i]; !ok && !dfs(i, 0) {
			return false
		}
	}
	return true
}
