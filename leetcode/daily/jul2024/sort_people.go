package jul2024

import "slices"

// https://leetcode.com/problems/sort-the-people/description/
func sortPeople(names []string, heights []int) []string {
	type Person struct {
		name   string
		height int
	}
	people := make([]*Person, 0, len(names))
	for i := range names {
		people = append(people, &Person{names[i], heights[i]})
	}
	slices.SortFunc(people, func(a, b *Person) int {
		return b.height - a.height
	})

	sortedNameByHeight := make([]string, 0, len(names))
	for i := range people {
		sortedNameByHeight = append(sortedNameByHeight, people[i].name)
	}
	return sortedNameByHeight
}
