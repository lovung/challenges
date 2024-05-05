package apr2024

import (
	"maps"

	"github.com/lovung/ds/matrix"
)

type cellmap struct {
	p     point
	h, v  map[point]bool
	other map[point]bool
}

func (c *cellmap) getAll() map[point]bool {
	m := make(map[point]bool)
	for k := range c.h {
		m[k] = true
	}
	for k := range c.v {
		m[k] = true
	}
	for k := range c.other {
		m[k] = true
	}
	return m
}

func (c *cellmap) merge(a, b *cellmap) {
	if a != nil {
		c.h = maps.Clone(a.h)
	} else {
		c.h = make(map[point]bool)
	}
	if b != nil {
		c.v = maps.Clone(b.v)
	} else {
		c.v = make(map[point]bool)
	}
	c.h[c.p] = true
	c.v[c.p] = true
	if a != nil && b != nil {
		c.other = mapInsection(a.getAll(), b.getAll())
	}
}

// https://leetcode.com/problems/maximal-rectangle/
func maximalRectangle1(mat [][]byte) int {
	rect := make([][]*cellmap, len(mat))
	for i := range rect {
		rect[i] = make([]*cellmap, len(mat[i]))
	}
	res := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == '0' {
				continue
			}
			var a, b *cellmap
			if matrix.In(mat, i-1, j) {
				a = rect[i-1][j]
			}
			if matrix.In(mat, i, j-1) {
				b = rect[i][j-1]
			}
			rect[i][j] = &cellmap{
				p: point{i, j},
			}
			rect[i][j].merge(a, b)
			res = max(res, calMaxRectCellMap(rect[i][j]))
		}
	}
	return res
}

func calMaxRectCellMap(c *cellmap) int {
	res := 0
	for p := range c.h {
		res = max(res, (c.p.x-p.x+1)*(c.p.y-p.y+1))
	}
	for p := range c.v {
		res = max(res, (c.p.x-p.x+1)*(c.p.y-p.y+1))
	}
	for p := range c.other {
		res = max(res, (c.p.x-p.x+1)*(c.p.y-p.y+1))
	}
	return res
}

func mapInsection(mapA, mapB map[point]bool) map[point]bool {
	insct := make(map[point]bool)
	for k := range mapB {
		if mapA[k] {
			insct[k] = true
		}
	}
	return insct
}
