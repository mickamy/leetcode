package main

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, ch := range word {
		if _, exists := curr.children[ch]; !exists {
			curr.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		curr = curr.children[ch]
	}
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, ch := range word {
		if _, exists := curr.children[ch]; !exists {
			return false
		}
		curr = curr.children[ch]
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, ch := range prefix {
		if _, exists := curr.children[ch]; !exists {
			return false
		}
		curr = curr.children[ch]
	}
	return true
}
