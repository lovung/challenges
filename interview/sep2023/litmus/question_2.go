package litmus

func findAllHobbyists(hobby string, hobbies map[string][]string) []string {
	hobbyistMap := make(map[string][]string)
	for personName, hobbyList := range hobbies {
		for i := range hobbyList {
			hobbyistMap[hobbyList[i]] = append(hobbyistMap[hobbyList[i]], personName)
		}
	}
	return hobbyistMap[hobby]
}
