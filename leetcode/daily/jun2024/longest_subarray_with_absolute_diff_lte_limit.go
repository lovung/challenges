package jun2024

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit
func longestSubarray1(nums []int, limit int) int {
	res := 0
	cnt := make(map[int]int)
	minVal, maxVal := int(1e9), 1
	for l, r := 0, 0; r < len(nums); r++ {
		minVal = min(minVal, nums[r])
		maxVal = max(maxVal, nums[r])
		cnt[nums[r]]++
		for l <= r && maxVal-minVal > limit {
			cnt[nums[l]]--
			if cnt[nums[l]] == 0 {
				delete(cnt, nums[l])
				if nums[l] == maxVal {
					maxVal = 1
					for k := range cnt {
						maxVal = max(maxVal, k)
					}
				} else if nums[l] == minVal {
					minVal = int(1e9)
					for k := range cnt {
						minVal = min(minVal, k)
					}
				}
			}
			l++
		}
		res = max(res, r-l+1)
		if res == len(nums) {
			return res
		}
	}
	return res
}

// Dequeue
func longestSubarray2(nums []int, limit int) int {
	var minD, maxD dequeue[int]
	l, res := 0, 1
	for r, nr := range nums {
		for minD.Len() > 0 && minD.Back() > nr {
			minD.PopBack()
		}
		for maxD.Len() > 0 && maxD.Back() < nr {
			maxD.PopBack()
		}
		minD.PushBack(nr)
		maxD.PushBack(nr)
		for ; maxD.Front()-minD.Front() > limit; l++ {
			if maxD.Front() == nums[l] {
				maxD.PopFront()
			}
			if minD.Front() == nums[l] {
				minD.PopFront()
			}
		}
		res = max(res, r-l+1)
	}
	return res
}

type dequeue[T any] struct {
	arr []T
}

func (d *dequeue[T]) Len() int {
	return len(d.arr)
}

func (d *dequeue[T]) Back() T {
	return d.arr[len(d.arr)-1]
}

func (d *dequeue[T]) Front() T {
	return d.arr[0]
}

func (d *dequeue[T]) PopBack() {
	d.arr = d.arr[:len(d.arr)-1]
}

func (d *dequeue[T]) PopFront() {
	d.arr = d.arr[1:]
}

func (d *dequeue[T]) PushBack(v T) {
	d.arr = append(d.arr, v)
}
