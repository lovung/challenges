package contest

import "sort"

func sortPeople(names []string, heights []int) []string {
	type person struct {
		name   string
		height int
	}
	people := make([]*person, len(names))
	for i := range names {
		people[i] = &person{
			name:   names[i],
			height: heights[i],
		}
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].height > people[j].height
	})
	sortNames := make([]string, len(names))
	for i := range people {
		sortNames[i] = people[i].name
	}
	return sortNames
}
