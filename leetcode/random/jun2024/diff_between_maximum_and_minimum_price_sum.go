package jun2024

import (
	"slices"

	llq "github.com/emirpasic/gods/v2/queues/linkedlistqueue"
)

// https://leetcode.com/problems/difference-between-maximum-and-minimum-price-sum/
// TLE: 54 / 58 testcases passed
func maxOutput1(n int, edges [][]int, price []int) int64 {
	nodes := make([]*node, 0, n)
	for i := range n {
		nodes = append(nodes, &node{[]*priceNode{{int64(price[i]), int64(price[i])}}})
	}
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	leaves := llq.New[int]()
	for i, near := range adj {
		if len(near) == 1 {
			leaves.Enqueue(i)
		}
	}
	res := int64(0)
	for !leaves.Empty() {
		l, _ := leaves.Dequeue()
		if len(adj[l]) < 1 {
			continue
		}
		near := adj[l][0]
		adj[near] = slices.DeleteFunc(adj[near], func(i int) bool { return i == l })
		if len(adj[near]) == 1 {
			leaves.Enqueue(near)
		}
		res = max(res, nodes[l].mergeTo(nodes[near]))
	}
	return res
}

type node struct {
	n []*priceNode
}

func (n *node) mergeTo(nn *node) (res int64) {
	newNodes := make([]*priceNode, 0, len(n.n))
	for _, pn := range n.n {
		newNode, _res := nn.tryMerge(pn)
		newNodes = append(newNodes, newNode)
		res = max(res, _res)
	}
	nn.n = append(nn.n, newNodes...)
	nn.n = slices.DeleteFunc(nn.n, func(e *priceNode) bool { return e == nil })
	return res
}

func (n *node) tryMerge(pn *priceNode) (nn *priceNode, res int64) {
	for i, ni := range n.n {
		if ni == nil {
			continue
		}
		if ni.zero() {
			nn = ni.clone().merge(pn)
			res = max(res, nn.gap())
		} else {
			if ni.sureGt(nn) {
				nn = nil
			} else if nn.sureGt(ni) {
				(n.n)[i] = nil
			}
			res = max(res, sum(ni, pn))
		}
	}
	return nn, res
}

type priceNode struct {
	total, min int64
}

func (pn *priceNode) zero() bool {
	return pn.total == pn.min
}

func (pn *priceNode) gap() int64 {
	return pn.total - pn.min
}

func sum(a, b *priceNode) int64 {
	return a.total + b.total - min(a.min, b.min)
}

func (pn *priceNode) merge(other *priceNode) *priceNode {
	pn.total += other.total
	pn.min = other.min
	return pn
}

func (pn *priceNode) clone() *priceNode {
	return &priceNode{pn.total, pn.min}
}

func (pn *priceNode) sureGt(other *priceNode) bool {
	if other == nil {
		return true
	}
	if pn.gap() > other.gap() && pn.total > other.total {
		return true
	}
	if pn.min == other.min {
		return pn.total > other.total
	}
	return false
}
