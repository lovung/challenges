package jun2024

import (
	"strings"
)

// https://leetcode.com/problems/replace-words/
func replaceWords(dictionary []string, sentence string) string {
	t := NewSimpleTrie()
	for i := range dictionary {
		t.Insert([]byte(dictionary[i]))
	}
	words := strings.Split(sentence, " ")
	for i := range words {
		if l := t.RemainNotFound([]byte(words[i])); l < len(words[i]) && l > 0 {
			words[i] = words[i][:len(words[i])-l]
		}
	}
	return strings.Join(words, " ")
}

// SimpleTrie only support `a` to `z`
type SimpleTrie struct {
	end      bool
	children [26]*SimpleTrie
}

func NewSimpleTrie() SimpleTrie {
	return SimpleTrie{}
}

const a = 'a'

func (this *SimpleTrie) Insert(word []byte) {
	if len(word) == 0 {
		this.end = true
		return
	}
	if this.children[word[0]-a] == nil {
		this.children[word[0]-a] = &SimpleTrie{}
	}
	this.children[word[0]-a].Insert(word[1:])
}

func (this *SimpleTrie) RemainNotFound(word []byte) int {
	if this.end {
		return len(word)
	}
	if len(word) == 0 {
		return 0
	}
	if this.children[word[0]-a] == nil {
		return 0
	}
	return this.children[word[0]-a].RemainNotFound(word[1:])
}
