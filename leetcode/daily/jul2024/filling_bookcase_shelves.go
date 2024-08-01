package jul2024

import (
	"math"

	"github.com/emirpasic/gods/v2/queues/linkedlistqueue"
)

type book struct {
	thickness, height int
}

// https://leetcode.com/problems/filling-bookcase-shelves/description/
func minHeightShelves(books [][]int, shelfWidth int) int {
	bookList := make([]book, 0, len(books))
	maxHeight := 0
	for _, b := range books {
		bookList = append(bookList, book{b[0], b[1]})
		maxHeight = max(maxHeight, b[1])
	}
	type item struct {
		startBookIdx int
		curHeight    int
	}
	q := linkedlistqueue.New[item]()
	q.Enqueue(item{})
	res := math.MaxInt
	for !q.Empty() {
		qi, _ := q.Dequeue()
		if qi.startBookIdx == len(bookList) {
			res = min(res, qi.curHeight)
		}
		prevEndIdx := -1
		for h := 1; h <= maxHeight; h++ {
			endIdx := consume(bookList, qi.startBookIdx, h, shelfWidth)
			if endIdx > prevEndIdx && endIdx >= qi.startBookIdx {
				q.Enqueue(item{endIdx + 1, qi.curHeight + h})
				prevEndIdx = endIdx
			}
		}
	}
	return res
}

// return endIdx can consume with the height
func consume(bookList []book, startIdx, height, shelfWidth int) int {
	i := startIdx
	for ; i < len(bookList); i++ {
		if bookList[i].height > height {
			break
		}
		if bookList[i].thickness > shelfWidth {
			break
		}
		shelfWidth -= bookList[i].thickness
	}
	return i - 1
}

func minHeightShelves2(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		w, h := books[i-1][0], books[i-1][1]
		dp[i] = dp[i-1] + h
		for j := i - 1; j > 0; j-- {
			w += books[j-1][0]
			if w > shelfWidth {
				break
			}
			h = max(h, books[j-1][1])
			dp[i] = min(dp[i], dp[j-1]+h)
		}
	}
	return dp[n]
}
