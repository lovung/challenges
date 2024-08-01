package jul2024

func floydWarshall(graph [][]int, v int) [][]int {
	dist := make([][]int, v)
	for i := range v {
		dist[i] = make([]int, v)
		copy(dist[i], graph[i])
	}

	for k := range v {
		for i := range v {
			for j := range v {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return dist
}

func minimumCost(
	source string, target string,
	original []byte, changed []byte,
	cost []int,
) int64 {
	const INF = 1e8
	const num = 26
	g := make([][]int, num)
	for i := range num {
		g[i] = make([]int, num)
	}
	for i := range num {
		for j := range num {
			if i == j {
				g[i][j] = 0
			} else {
				g[i][j] = INF
			}
		}
	}

	const a = 'a'
	for i := range original {
		tmp := original[i] - a
		tmp2 := changed[i] - a
		g[tmp][tmp2] = min(g[tmp][tmp2], cost[i])
	}

	dist := floydWarshall(g, num)
	ans := int64(0)
	for i, s := range source {
		t := target[i]
		if dist[s-a][t-a] == INF {
			return -1
		}
		ans += int64(dist[s-a][t-a])
	}
	return ans
}
