package apr2024

// https://leetcode.com/problems/number-of-wonderful-substrings/
func wonderfulSubstrings(word string) int64 {
	count := [1024]int64{}      // cnt[state] stores how many times the state occurs
	count[0] = 1                // empty string gives case where all characters occur even number of times
	res, bitmask := int64(0), 0 // bitmask: current state
	for i := range word {
		idx := int(word[i] - a)
		bitmask ^= 1 << idx   // update state
		res += count[bitmask] // add count of same previous states
		for j := range 10 {
			k := bitmask ^ (1 << j) // try flick each switch
			res += count[k]
		}
		count[bitmask]++ // add 1 to count of times we've seen current state
	}
	return res
}

func wonderfulSubstrings1(word string) int64 {
	cter := counter{}
	res := int64(len(word))
	for i := range word {
		cter.add(word[i])
		if i > 0 {
			cter.remove(word[i-1])
		}
		newCter := cter
		for j := i + 1; j < len(word); j++ {
			newCter.add(word[j])
			if newCter.isWonderful() {
				res++
			}
		}
	}
	return res
}

type counter struct {
	cnt [10]int
	odd int
}

func (c *counter) isWonderful() bool {
	return c.odd <= 1
}

func (c *counter) add(b byte) {
	c.cnt[b-a]++
	if c.cnt[b-a]&1 == 0 {
		c.odd--
	} else {
		c.odd++
	}
}

func (c *counter) remove(b byte) {
	c.cnt[b-a]--
	if c.cnt[b-a]&1 == 0 {
		c.odd--
	} else {
		c.odd++
	}
}
