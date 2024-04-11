package apr2024

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode.com/problems/copy-list-with-random-pointer/description/
func copyRandomList(head *Node) *Node {
	cloneMap := make(map[*Node]*Node)
	var prevNode *Node
	for n := head; n != nil; n = n.Next {
		newNode := &Node{
			Val: n.Val,
		}
		if prevNode != nil {
			prevNode.Next = newNode
		}
		cloneMap[n] = newNode
		prevNode = newNode
	}
	for n := head; n != nil; n = n.Next {
		cloneMap[n].Random = cloneMap[n.Random]
	}
	return cloneMap[head]
}
