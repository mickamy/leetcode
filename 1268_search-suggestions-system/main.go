package main

import "slices"

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := NewTrie()
	for _, product := range products {
		trie.Insert(product)
	}

	ans := make([][]string, len(searchWord))
	for i := range searchWord {
		candidates := trie.SearchWordsByPrefix(searchWord[:i+1])
		slices.Sort(candidates)
		if len(candidates) > 3 {
			candidates = candidates[:3]
		}
		ans[i] = candidates
	}
	return ans
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
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

func (t *Trie) SearchWordsByPrefix(prefix string) []string {
	curr := t.root
	for _, ch := range prefix {
		if _, exists := curr.children[ch]; !exists {
			return []string{}
		}
		curr = curr.children[ch]
	}

	var candidates []string
	var dfs func(node *TrieNode, curr string)
	dfs = func(node *TrieNode, curr string) {
		if node.isEnd {
			candidates = append(candidates, curr)
		}

		for ch, child := range node.children {
			dfs(child, curr+string(ch))
		}
	}

	dfs(curr, prefix)
	return candidates
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
