package biweekly132

import "golang.org/x/exp/constraints"

// https://leetcode.com/problems/find-the-first-player-to-win-k-games-in-a-row/description/
// TLE
func findWinningPlayer1(skills []int, k int) int {
	if k >= len(skills) {
		return FirstMaxValueIndex(skills)
	}
	type player struct{ idx, skill, win int }
	players := make([]*player, len(skills))
	for i := range skills {
		players[i] = &player{i, skills[i], 0}
	}
	for {
		if players[0].skill < players[1].skill {
			players[0], players[1] = players[1], players[0]
		}
		players[0].win++
		if players[0].win >= k {
			return players[0].idx
		}
		tmp := players[1]
		copy(players[1:], players[2:])
		players[len(players)-1] = tmp
	}
}

func FirstMaxValueIndex[T constraints.Ordered](s []T) int {
	maxVal := s[0]
	maxIdx := -1
	for i := range s {
		if maxVal < s[i] {
			maxVal = s[i]
			maxIdx = i
		}
	}
	return maxIdx
}

func findWinningPlayer(skills []int, k int) int {
	if k >= len(skills) {
		return FirstMaxValueIndex(skills)
	}
	type player struct {
		idx, skill, win int
		next            *player
	}
	var head, tail *player
	for i := range skills {
		if head == nil {
			head = &player{i, skills[i], 0, nil}
			tail = head
		} else {
			tail.next = &player{i, skills[i], 0, nil}
			tail = tail.next
		}
	}
	for {
		if head.skill < head.next.skill {
			head, tail.next, head.next = head.next, head, nil
		} else {
			head.next, tail.next, head.next.next = head.next.next, head.next, nil
		}
		tail = tail.next
		head.win++
		if head.win >= k {
			return head.idx
		}
	}
}
