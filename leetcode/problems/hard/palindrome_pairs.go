package hard

type Trie struct {
	next  [26]*Trie
	list  []int
	index int
}

// Link: https://leetcode.com/problems/palindrome-pairs/
func palindromePairs(words []string) [][]int {
	rootTrie := &Trie{
		index: -1,
		list:  make([]int, 0),
	}
	for i := range words {
		addWord(rootTrie, words[i], i)
	}

	res := make([][]int, 0)
	for i := range words {
		search(words, i, rootTrie, &res)
	}
	return res
}

func search(words []string, i int, root *Trie, res *[][]int) {
	for j, r := range words[i] {
		if root.index >= 0 &&
			root.index != i &&
			isPalindrome(words[i], j, len(words[i])-1) {
			*res = append(*res, []int{i, root.index})
		}
		root = root.next[r-a]
		if root == nil {
			return
		}
	}
	for _, j := range root.list {
		if i == j {
			continue
		}
		*res = append(*res, []int{i, j})
	}
}

func addWord(root *Trie, word string, index int) {
	for i := len(word) - 1; i >= 0; i-- {
		j := word[i] - a
		if root.next[j] == nil {
			root.next[j] = &Trie{
				index: -1,
				list:  make([]int, 0),
			}
		}
		if isPalindrome(word, 0, i) {
			root.list = append(root.list, index)
		}
		root = root.next[j]
	}
	root.list = append(root.list, index)
	root.index = index
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
