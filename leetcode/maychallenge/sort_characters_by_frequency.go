package maychallenge

// Link: https://leetcode.com/problems/sort-characters-by-frequency/
func frequencySort(s string) string {
	n := len(s)
	arr := make([]Item, 256)
	for _, e := range s {
		arr[e].count++
		arr[e].char = e
	}
	quickSort(arr, 0, 255)
	out := make([]rune, n)
	k := 0
	for i := 255; i >= 0; i-- {
		if arr[i].count == 0 {
			break
		}
		for j := 0; j < arr[i].count; j++ {
			out[k] = arr[i].char
			k++
		}
	}
	return string(out)
}

// Item to save the count
type Item struct {
	count int
	char  rune
}

func quickSort(arr []Item, l, r int) {
	if l < r {
		pi := partition(arr, l, r)

		quickSort(arr, l, pi-1)
		quickSort(arr, pi+1, r)
	}
}

func partition(arr []Item, l, r int) int {
	pivot := arr[r].count

	i := l - 1
	for j := l; j <= r; j++ {
		if arr[j].count < pivot {
			i++
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
		}
	}
	tmp := arr[i+1]
	arr[i+1] = arr[r]
	arr[r] = tmp
	return i + 1
}
