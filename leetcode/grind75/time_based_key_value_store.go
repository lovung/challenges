package grind75

// Link: https://leetcode.com/problems/time-based-key-value-store/
type TimeMap struct {
	store map[string][]*timeItem
}

type timeItem struct {
	timestamp int
	value     string
}

func NewTimeMap() TimeMap {
	return TimeMap{
		store: make(map[string][]*timeItem),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], &timeItem{
		timestamp: timestamp,
		value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.store[key]) == 0 {
		return ""
	}
	return this.binarySearch(this.store[key], timestamp)
}

func (this *TimeMap) binarySearch(items []*timeItem, timestamp int) string {
	if timestamp < items[0].timestamp {
		return ""
	}
	l, r := 0, len(items)-1
	for l <= r {
		mid := (l + r) / 2
		if items[mid].timestamp == timestamp {
			return items[mid].value
		}
		if items[mid].timestamp > timestamp {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return items[r].value
}
