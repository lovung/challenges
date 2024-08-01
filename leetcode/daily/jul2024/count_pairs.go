package jul2024

import "github.com/lovung/ds/trees"

// https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/description/
func countPairs(root *trees.TreeNode[int], distance int) int {
	var (
		res     int
		leafs   []*trees.TreeNode[int]
		parents = make(map[*trees.TreeNode[int]]*trees.TreeNode[int])
	)
	loadLeafs(root, &leafs, parents)
	visited := make(map[*trees.TreeNode[int]]bool)
	for _, l := range leafs {
		if parents[l] != nil {
			visited[l] = true
			dfs(parents[l], parents, visited, distance-1, &res)
			clear(visited)
		}
	}
	return res / 2
}

func isLeaf(node *trees.TreeNode[int]) bool {
	return node != nil && node.Right == nil && node.Left == nil
}

func dfs(
	node *trees.TreeNode[int], parents map[*trees.TreeNode[int]]*trees.TreeNode[int],
	visited map[*trees.TreeNode[int]]bool,
	remainStep int, cnt *int,
) {
	if node == nil {
		return
	}
	if remainStep < 0 {
		return
	}
	visited[node] = true
	if isLeaf(node) {
		*cnt++
		return
	}
	if node.Left != nil && !visited[node.Left] {
		dfs(node.Left, parents, visited, remainStep-1, cnt)
	}
	if node.Right != nil && !visited[node.Right] {
		dfs(node.Right, parents, visited, remainStep-1, cnt)
	}
	if !visited[parents[node]] {
		dfs(parents[node], parents, visited, remainStep-1, cnt)
	}
}

func loadLeafs(root *trees.TreeNode[int], leafs *[]*trees.TreeNode[int], parents map[*trees.TreeNode[int]]*trees.TreeNode[int]) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		*leafs = append(*leafs, root)
		return
	}
	if root.Right != nil {
		parents[root.Right] = root
	}
	if root.Left != nil {
		parents[root.Left] = root
	}
	loadLeafs(root.Left, leafs, parents)
	loadLeafs(root.Right, leafs, parents)
}

func countPairs2(root *trees.TreeNode[int], distance int) int {
	var run func(*trees.TreeNode[int]) ([]int, int)
	run = func(node *trees.TreeNode[int]) ([]int, int) {
		if node == nil {
			return []int{}, 0
		}
		if node.Left == nil && node.Right == nil {
			return []int{1}, 0
		}
		lLeaves, lCount := run(node.Left)
		rLeaves, rCount := run(node.Right)
		count := lCount + rCount
		for _, left := range lLeaves {
			for _, right := range rLeaves {
				if left+right <= distance {
					count++
				}
			}
		}
		leaves := make([]int, 0, len(lLeaves)+len(rLeaves))
		for _, leaf := range lLeaves {
			if leaf < 9 {
				leaves = append(leaves, leaf+1)
			}
		}
		for _, leaf := range rLeaves {
			if leaf < 9 {
				leaves = append(leaves, leaf+1)
			}
		}
		return leaves, count
	}

	_, result := run(root)
	return result
}
