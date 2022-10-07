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

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.end = true
		return
	}
	if this.children[word[0]-a] == nil {
		this.children[word[0]-a] = &Trie{}
	}
	this.children[word[0]-a].Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.end
	}
	if this.children[word[0]-a] == nil {
		return false
	}
	return this.children[word[0]-a].Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if this.children[prefix[0]-a] == nil {
		return false
	}
	return this.children[prefix[0]-a].StartsWith(prefix[1:])
}
