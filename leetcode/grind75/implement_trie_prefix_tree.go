package grind75

// Link: https://leetcode.com/problems/implement-trie-prefix-tree/
type Trie struct {
	end      bool
	children [26]*Trie
}

func NewTrie() Trie {
	return Trie{}
}

const a = 'a'

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		t.end = true
		return
	}
	if t.children[word[0]-a] == nil {
		t.children[word[0]-a] = &Trie{}
	}
	t.children[word[0]-a].Insert(word[1:])
}

func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return t.end
	}
	if t.children[word[0]-a] == nil {
		return false
	}
	return t.children[word[0]-a].Search(word[1:])
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if t.children[prefix[0]-a] == nil {
		return false
	}
	return t.children[prefix[0]-a].StartsWith(prefix[1:])
}
