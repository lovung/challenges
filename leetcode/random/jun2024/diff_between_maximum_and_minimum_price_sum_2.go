package jun2024

// https://leetcode.com/problems/difference-between-maximum-and-minimum-price-sum/
func maxOutput(n int, edges [][]int, price []int) int64 {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	var res int64
	dfs(adj, price, &res, 0, -1)
	return res
}

func dfs(graph [][]int, price []int, res *int64, cur, pre int) []int64 {
	curMax := []int64{int64(price[cur]), 0}
	for _, child := range graph[cur] {
		if child == pre {
			continue
		}
		sub := dfs(graph, price, res, child, cur)
		*res = max(*res, curMax[0]+sub[1])
		*res = max(*res, curMax[1]+sub[0])
		curMax[0] = max(curMax[0], sub[0]+int64(price[cur]))
		curMax[1] = max(curMax[1], sub[1]+int64(price[cur]))
	}
	return curMax
}
