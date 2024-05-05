package oct2022

import (
	"sort"
)

// Link: https://leetcode.com/problems/my-calendar-iii/
type MyCalendarThree struct {
	events []*calendarEvent
}

type calendarEvent struct {
	at  int
	cnt int
}

func NewMyCalendarThree() MyCalendarThree {
	return MyCalendarThree{
		events: make([]*calendarEvent, 0),
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	this.events = append(this.events, &calendarEvent{
		at:  start,
		cnt: 1,
	}, &calendarEvent{
		at:  end,
		cnt: -1,
	})
	sort.Slice(this.events, func(i, j int) bool {
		if this.events[i].at != this.events[j].at {
			return this.events[i].at < this.events[j].at
		}
		return this.events[i].cnt < this.events[j].cnt
	})
	maxVal := 0
	cnt := 0
	for _, e := range this.events {
		cnt += e.cnt
		maxVal = max(maxVal, cnt)
	}
	return maxVal
}
