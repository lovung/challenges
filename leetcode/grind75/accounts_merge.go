package grind75

import (
	"sort"

	"github.com/lovung/ds/graphs"
)

// Link: https://leetcode.com/problems/accounts-merge/
func accountsMerge(accounts [][]string) [][]string {
	uf := graphs.NewUnionFind(len(accounts))
	emailMap := make(map[string]int)
	for i := range accounts {
		for j := 1; j < len(accounts[i]); j++ {
			if idx, ok := emailMap[accounts[i][j]]; ok {
				uf.Union(idx, i)
			} else {
				emailMap[accounts[i][j]] = i
			}
		}
	}

	mergedAccounts := make([][]string, len(accounts))
	for i := range accounts {
		groupIdx := uf.Find(i)
		putEmails(&mergedAccounts[groupIdx], accounts[i])
	}
	return removeEmptyAccounts(mergedAccounts)
}

func removeEmptyAccounts(accounts [][]string) [][]string {
	res := make([][]string, 0, len(accounts))
	for i := range accounts {
		if len(accounts[i]) > 0 {
			sort.Strings(accounts[i][1:])
			res = append(res, accounts[i])
		}
	}
	return res
}

func putEmails(accounts *[]string, newEmails []string) {
	if len(*accounts) == 0 {
		*accounts = make([]string, 0, len(newEmails))
		*accounts = append(*accounts, newEmails[0])
	}
	emailMap := make(map[string]struct{})
	for i := 1; i < len(*accounts); i++ {
		emailMap[(*accounts)[i]] = struct{}{}
	}
	for i := 1; i < len(newEmails); i++ {
		if _, ok := emailMap[newEmails[i]]; !ok {
			*accounts = append(*accounts, newEmails[i])
			emailMap[newEmails[i]] = struct{}{}
		}
	}
}
