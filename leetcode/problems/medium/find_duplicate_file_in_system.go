package medium

import "strings"

// Link: https://leetcode.com/problems/find-duplicate-file-in-system/
func findDuplicate(paths []string) [][]string {
	contentToFileNames := make(map[string][]string)
	for i := range paths {
		pathToMapByContent(paths[i], contentToFileNames)
	}
	res := make([][]string, 0)
	for _, v := range contentToFileNames {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}

func pathToMapByContent(path string, contentToFileNames map[string][]string) {

	pathInfo := strings.Split(path, " ")
	dir := pathInfo[0]
	for i := 1; i < len(pathInfo); i++ {
		fileName, content := nameAndContent(pathInfo[i])
		contentToFileNames[content] = append(contentToFileNames[content], dir+"/"+fileName)
	}
}

func nameAndContent(full string) (string, string) {
	startedContextIdx := strings.Index(full, "(")
	return full[:startedContextIdx], full[startedContextIdx+1 : len(full)-1]
}
