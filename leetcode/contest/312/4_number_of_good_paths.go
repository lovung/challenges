package contest

import (
	"sort"

	"github.com/lovung/ds/graphs"
)

// Link: https://leetcode.com/problems/number-of-good-paths/
func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	uf := graphs.NewUnionFind(n)
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	added := make([]bool, n)

	kvs := sortedNodeByValues(vals)
	res := 0
	for _, kvi := range kvs {
		v := kvi.v
		for _, i := range v {
			added[i] = true
			for _, j := range adj[i] {
				if added[j] {
					uf.Union(i, j)
				}
			}
		}
		cnt := make(map[int]int)
		for _, i := range v {
			cnt[uf.Find(i)]++
		}
		for _, v := range cnt {
			res += v * (v + 1) / 2
		}
	}
	return res
}

type kv struct {
	k int
	v []int
}

func sortedNodeByValues(vals []int) []*kv {
	d := make(map[int][]int)
	for i := range vals {
		d[vals[i]] = append(d[vals[i]], i)
	}

	kvs := make([]*kv, 0, len(d))
	for k, v := range d {
		kvs = append(kvs, &kv{k, v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].k < kvs[j].k
	})
	return kvs
}
