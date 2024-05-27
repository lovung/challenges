package weekly399

// Thanks to @jeffreyhu8
// https://leetcode.com/problems/maximum-sum-of-subsequence-with-non-adjacent-elements/solutions/5208917/segment-tree

const (
	inf = 1e18
	mod = 1e9 + 7
)

// Node represents a node in the Lazy Segment Tree
type Node struct {
	l, r   *Node
	lo, hi uint
	// selected stores the maximum subarray sums for different conditions
	selected [2][2]int
}

// NewNode creates a new Node
func NewNode(nums []int, lo, hi uint) *Node {
	n := &Node{lo: lo, hi: hi}
	if lo < hi {
		mid := (lo + hi) / 2
		n.l = NewNode(nums, lo, mid)
		n.r = NewNode(nums, mid+1, hi)
		n.combine()
	} else {
		n.selected[0][0] = 0
		n.selected[0][1] = -inf
		n.selected[1][0] = -inf
		n.selected[1][1] = nums[lo]
	}
	return n
}

// combine combines the information from child nodes
func (n *Node) combine() {
	n.selected[0][0] = max(
		n.l.selected[0][0]+n.r.selected[0][0],
		n.l.selected[0][1]+n.r.selected[0][0],
		n.l.selected[0][0]+n.r.selected[1][0],
	)
	n.selected[0][1] = max(
		n.l.selected[0][0]+n.r.selected[0][1],
		n.l.selected[0][1]+n.r.selected[0][1],
		n.l.selected[0][0]+n.r.selected[1][1],
	)
	n.selected[1][0] = max(
		n.l.selected[1][0]+n.r.selected[0][0],
		n.l.selected[1][1]+n.r.selected[0][0],
		n.l.selected[1][0]+n.r.selected[1][0],
	)
	n.selected[1][1] = max(
		n.l.selected[1][0]+n.r.selected[0][1],
		n.l.selected[1][1]+n.r.selected[0][1],
		n.l.selected[1][0]+n.r.selected[1][1],
	)
}

// update updates the value at a specific index
func (n *Node) update(i int, x int) {
	if uint(i) < n.lo || n.hi < uint(i) {
		return
	}
	if n.lo == n.hi {
		n.selected[0][0] = 0
		n.selected[1][1] = x
		return
	}
	n.l.update(i, x)
	n.r.update(i, x)
	n.combine()
}

func maximumSumSubsequence(nums []int, queries [][]int) int {
	segmentTree := NewNode(nums, 0, uint(len(nums)-1))
	var res int

	for _, query := range queries {
		segmentTree.update(query[0], query[1])
		res += max(segmentTree.selected[0][0],
			segmentTree.selected[0][1],
			segmentTree.selected[1][0],
			segmentTree.selected[1][1])
		res %= mod
	}

	return res
}
