package apr2024

import "github.com/lovung/ds/matrix"

type point struct {
	x, y int
}

type rect struct {
	p1, p2 point
}

func (r rect) cover(p point) bool {
	return r.p1.x <= p.x && p.x <= r.p2.x &&
		r.p1.y <= p.y && p.y <= r.p2.y
}

func (r rect) area() int {
	return (r.p2.x - r.p1.x + 1) * (r.p2.y - r.p1.y + 1)
}

func (r rect) valid() bool {
	return r.p1.x <= r.p2.x && r.p1.y <= r.p2.y
}

func (r rect) merge(other rect) rect {
	return rect{
		p1: point{
			x: max(r.p1.x, other.p1.x),
			y: max(r.p1.y, other.p1.y),
		},
		p2: point{
			x: min(r.p2.x, other.p2.x),
			y: min(r.p2.y, other.p2.y),
		},
	}
}

type cell struct {
	p     point
	rects []rect
}

func (c *cell) add(newPoint point) {
	for _, r := range c.rects {
		if r.cover(newPoint) {
			return
		}
	}
	c.rects = append(c.rects, rect{newPoint, c.p})
}

func (c *cell) merge(a, b *cell) {
	if a == nil && b == nil {
		c.add(c.p)
		return
	}
	if a != nil && b != nil {
		for _, ar := range a.rects {
			for _, br := range b.rects {
				cr := ar.merge(br)
				if cr.valid() {
					c.add(cr.p1)
				}
			}
		}
	}
	if b != nil {
		for _, br := range b.rects {
			cr := br.merge(rect{point{c.p.x, 0}, point{c.p.x, c.p.y}})
			if cr.valid() {
				c.add(cr.p1)
			}
		}
	}
	if a != nil {
		for _, ar := range a.rects {
			cr := ar.merge(rect{point{0, c.p.y}, point{c.p.x, c.p.y}})
			if cr.valid() {
				c.add(cr.p1)
			}
		}
	}
}

func (c *cell) maxRectArea() int {
	res := 0
	for _, r := range c.rects {
		res = max(res, r.area())
	}
	return res
}

// https://leetcode.com/problems/maximal-rectangle/
func maximalRectangle2(mat [][]byte) int {
	rect := make([][]*cell, len(mat))
	for i := range rect {
		rect[i] = make([]*cell, len(mat[i]))
	}
	res := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == '0' {
				continue
			}
			var a, b *cell
			if matrix.In(mat, i-1, j) {
				a = rect[i-1][j]
			}
			if matrix.In(mat, i, j-1) {
				b = rect[i][j-1]
			}
			rect[i][j] = &cell{
				p: point{i, j},
			}
			rect[i][j].merge(a, b)
			res = max(res, rect[i][j].maxRectArea())
		}
	}
	return res
}
