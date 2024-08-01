package jul2024

import (
	"slices"

	"github.com/emirpasic/gods/v2/stacks/linkedliststack"
)

// https://leetcode.com/problems/robot-collisions/description/
func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	type robot struct {
		index, initPos, health int
		dir                    byte
	}
	robots := make([]*robot, len(positions))
	for i := range positions {
		robots[i] = &robot{
			index:   i,
			initPos: positions[i],
			health:  healths[i],
			dir:     directions[i],
		}
	}
	slices.SortFunc(robots, func(a, b *robot) int {
		return a.initPos - b.initPos
	})
	rStack := linkedliststack.New[*robot]()
	for _, robot := range robots {
		if robot.dir == 'R' {
			rStack.Push(robot)
		} else {
			for right, ok := rStack.Peek(); ok; right, ok = rStack.Peek() {
				if right.health == robot.health {
					right.health = 0
					robot.health = 0
					rStack.Pop()
					break
				} else if right.health > robot.health {
					right.health--
					robot.health = 0
					break
				} else {
					right.health = 0
					robot.health--
					rStack.Pop()
				}
			}
		}
	}
	robots = slices.DeleteFunc(robots, func(r *robot) bool {
		return r.health == 0
	})
	slices.SortFunc(robots, func(a, b *robot) int {
		return a.index - b.index
	})
	res := make([]int, len(robots))
	for i, r := range robots {
		res[i] = r.health
	}
	return res
}
