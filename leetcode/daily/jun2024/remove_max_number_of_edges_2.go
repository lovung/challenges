package jun2024

// https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/
// Disjoin set
func maxNumEdgesToRemove(n int, edges [][]int) int {
	ufA := NewUnionFind(n)
	ufB := NewUnionFind(n)
	// process type 3 first
	res := 0
	for _, e := range edges {
		if e[0] == 3 {
			if !ufA.Union(e[1]-1, e[2]-1) {
				res++
			}
			ufB.Union(e[1]-1, e[2]-1)
		}
	}
	for _, e := range edges {
		if e[0] == 1 {
			if !ufA.Union(e[1]-1, e[2]-1) {
				res++
			}
		} else if e[0] == 2 {
			if !ufB.Union(e[1]-1, e[2]-1) {
				res++
			}
		}
	}
	if ufA.Count() != 1 || ufB.Count() != 1 {
		return -1
	}
	return res
}

// UnionFind implemented based on https://en.wikipedia.org/wiki/Disjoint-set_data_structure
type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

// NewUnionFind init the UnionFind with size
func NewUnionFind(size int) UnionFind {
	uf := UnionFind{
		parent: make([]int, size),
		rank:   make([]int, size),
		count:  size,
	}
	for i := 0; i < size; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Count() int {
	return uf.count
}

// Find operation follows the chain of parent pointers from a specified query node x
// until it reaches a root element. This root element represents the set to which x belongs
// and may be x itself. Find returns the root element it reaches.
//
// Tarjan and Van Leeuwen also developed one-pass Find algorithms that retain the same worst-case complexity
// but are more efficient in practice
func (uf *UnionFind) Find(x int) int {
	for uf.parent[x] != x {
		x, uf.parent[x] = uf.parent[x], uf.parent[uf.parent[x]]
	}
	return x
}

// Union replaces the set containing x and the set containing y with their union.
func (uf *UnionFind) Union(x, y int) bool {
	x = uf.Find(x)
	y = uf.Find(y)

	if x == y {
		// x and y are already in the same set
		return false
	}

	// If necessary, rename variables to ensure that
	// x has rank at least as large as that of y
	if uf.rank[x] < uf.rank[y] {
		x, y = y, x
	}

	// Make x the new root
	uf.parent[y] = x
	// If necessary, increase the rank of x
	if uf.rank[x] == uf.rank[y] {
		uf.rank[x]++
	}
	uf.count--
	return true
}
