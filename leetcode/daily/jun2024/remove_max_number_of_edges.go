package jun2024

func maxNumEdgesToRemove_TLE(n int, edges [][]int) int {
	adjA := make([]map[int]struct{}, n)
	adjB := make([]map[int]struct{}, n)
	// process type 3 first
	res := 0
	for _, e := range edges {
		if e[0] == 3 {
			if alreadyConnected(adjA, e[1], e[2]) &&
				alreadyConnected(adjB, e[1], e[2]) {
				res++
				continue
			}
			addEdge(adjA, e[1], e[2])
			addEdge(adjB, e[1], e[2])
		}
	}
	for _, e := range edges {
		if e[0] == 1 {
			if alreadyConnected(adjA, e[1], e[2]) {
				res++
				continue
			}
			addEdge(adjA, e[1], e[2])
		} else if e[0] == 2 {
			if alreadyConnected(adjB, e[1], e[2]) {
				res++
				continue
			}
			addEdge(adjB, e[1], e[2])
		}
	}
	if cntEdge(adjA) < n-1 {
		return -1
	}
	if cntEdge(adjB) < n-1 {
		return -1
	}
	return res
}

func cntEdge(adj []map[int]struct{}) int {
	cnt := 0
	for i := range adj {
		cnt += len(adj[i])
	}
	return cnt / 2
}

func addEdge(adj []map[int]struct{}, a, b int) {
	if adj[a-1] == nil {
		adj[a-1] = make(map[int]struct{})
	}
	adj[a-1][b] = struct{}{}
	if adj[b-1] == nil {
		adj[b-1] = make(map[int]struct{})
	}
	adj[b-1][a] = struct{}{}
}

func alreadyConnected(adj []map[int]struct{}, from, to int) bool {
	if from == to {
		return true
	}
	return alreadyConnectedRecursive(adj, from, to, -1)
}

func alreadyConnectedRecursive(adj []map[int]struct{}, from, to int, prev int) bool {
	if from == to {
		return true
	}
	for k := range adj[from-1] {
		if k == prev {
			continue
		}
		if alreadyConnectedRecursive(adj, k, to, from) {
			return true
		}
	}
	return false
}
