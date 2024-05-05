package apr2024

// https://leetcode.com/problems/sum-of-distances-in-tree/
func sumOfDistancesInTree3(n int, edges [][]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	cnt := make([]int, n)
	res := make([]int, n)
	for i := range cnt {
		cnt[i] = 1
	}

	dfsCalculate(adj, cnt, res, 0, -1)
	dfsAdjust(adj, cnt, res, n, 0, -1)
	return res
}

func dfsCalculate(adj [][]int, cnt, res []int, node, parent int) {
	for _, child := range adj[node] {
		if child != parent {
			dfsCalculate(adj, cnt, res, child, node)
			cnt[node] += cnt[child]
			res[node] += res[child] + cnt[child]
		}
	}
}

func dfsAdjust(adj [][]int, cnt, res []int, n, node, parent int) {
	for _, child := range adj[node] {
		if child != parent {
			res[child] = res[node] - cnt[child] + (n - cnt[child])
			dfsAdjust(adj, cnt, res, n, child, node)
		}
	}
}
