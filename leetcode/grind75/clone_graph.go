package grind75

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	created := make(map[*Node]*Node)
	res := &Node{}
	dfsCloneGraph(&res, node, created)
	return res
}

func dfsCloneGraph(dstNode **Node, srcNode *Node, created map[*Node]*Node) {
	if srcNode == nil {
		*dstNode = nil
		return
	}
	if v, ok := created[srcNode]; ok {
		*dstNode = v
		return
	}
	*dstNode = &Node{
		Val:       srcNode.Val,
		Neighbors: make([]*Node, len(srcNode.Neighbors)),
	}
	created[srcNode] = *dstNode
	for i := range srcNode.Neighbors {
		dfsCloneGraph(&((*dstNode).Neighbors[i]), srcNode.Neighbors[i], created)
	}
}

func compareNodes(n1, n2 *Node) bool {
	visited := make(map[*Node]bool)
	return dfsCompareNodes(n1, n2, visited)
}

func dfsCompareNodes(n1, n2 *Node, visited map[*Node]bool) bool {
	if n1 == n2 {
		return false
	}
	if visited[n1] {
		return true
	}
	visited[n1] = true
	res := true
	for i := range n1.Neighbors {
		res = res && dfsCompareNodes(n1.Neighbors[i], n2.Neighbors[i], visited)
	}
	return res
}
