package apr2024

import (
	"sort"

	"github.com/lovung/ds/slice"
)

// https://leetcode.com/problems/reveal-cards-in-increasing-order/description/
func deckRevealedIncreasing(deck []int) []int {
	sort.Slice(deck, func(i, j int) bool { return deck[i] > deck[j] })
	curDeck := make([]int, 0, len(deck))
	for i := range deck {
		curDeck = slice.PushFront(curDeck, deck[i])
		if i < len(deck)-1 {
			slice.RotateRight(curDeck, 1)
		}
	}
	return curDeck
}
