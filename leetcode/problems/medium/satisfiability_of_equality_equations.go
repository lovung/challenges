package medium

import "github.com/lovung/ds/graphs"

// Link: https://leetcode.com/problems/satisfiability-of-equality-equations/
func equationsPossible(equations []string) bool {
	const a = 'a'
	equalUF := graphs.NewUnionFind(26)

	for _, e := range equations {
		if e[1] == '=' {
			equalUF.Union(int(e[0]-a), int(e[3]-a))
		}
	}

	for _, e := range equations {
		if e[1] == '!' && equalUF.Find(int(e[0]-a)) == equalUF.Find(int(e[3]-a)) {
			return false
		}
	}
	return true
}
