package weekly404

import "github.com/emirpasic/gods/v2/queues/linkedlistqueue"

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	d1 := diameter(edges1)
	d2 := diameter(edges2)
	return max(d1, d2, (d1+1)/2+(d2+1)/2+1)
}

func diameter(edges [][]int) int {
	m := len(edges) + 1
	res := 0
	al := make([][]int, m)
	degree := make([]int, m)
	depth := make([]int, m)
	vis := make([]bool, m)
	q := linkedlistqueue.New[int]()

	for _, e := range edges {
		al[e[0]] = append(al[e[0]], e[1])
		al[e[1]] = append(al[e[1]], e[0])
	}

	for i := 0; i < m; i++ {
		degree[i] = len(al[i])
		if degree[i] == 1 {
			q.Enqueue(i)
		}
	}

	for q.Size() > 0 {
		i, _ := q.Dequeue()
		vis[i] = true
		for _, j := range al[i] {
			degree[j]--
			if degree[j] == 1 {
				q.Enqueue(j)
			}
			if !vis[j] {
				res = max(res, depth[j]+depth[i]+1)
				depth[j] = max(depth[j], depth[i]+1)
			}
		}
	}

	return res
}
