package trie

import "strings"

// A trie is a tree-like data structure whose nodes store the letters of an alphabet. By structuring the nodes in
// a particular way, words and strings can be retrieved from the structure by traversing down a branch path of the tree.
// O(m)

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

func (t *Trie) Insert(w string) {
	w = strings.ToLower(w)
	length := len(w)
	currentNode := t.root
	for i := 0; i < length; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	w = strings.ToLower(w)
	length := len(w)
	currentNode := t.root
	for i := 0; i < length; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}
