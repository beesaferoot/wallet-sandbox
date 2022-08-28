package utils

import (
	"fmt"
)


/* 
ALGO:
 - support url path handler insert 
 - support wildcard url path handler

E.g:

/path
/path/*
/path/*\/specific

*/

type TrieNode struct {
	value rune
	handler Handler
	children  []*TrieNode
	isEndofPath bool 
	wildcard bool
	urlpath string 
	method string
}

type Trie struct {
	root *TrieNode
}


func (tn *TrieNode) Find(ch rune) *TrieNode {

	for _, wch := range tn.children {
		if wch.value == ch {
			return wch
		}
		if wch.wildcard {
			if ch == '/' {
				return wch.children[0]
			}
			return wch
		}
		
	}
	return nil
}

func (tn *TrieNode) Handler() Handler {
	return tn.handler
}

func (tn *TrieNode) Path() string {
	return tn.urlpath
}

func (tn *TrieNode) Method() string {
	return tn.method
}

func (t *Trie) AddHandler(path , method string ,  handler Handler) {
	currNode := t.root

	path_len := len(path)

	for i := 0; i < path_len; i++ {
		ch := path[i]
		node := currNode.Find(rune(ch))
		switch (ch) {
			case '*': {
				if node == nil {
					node = &TrieNode{value: rune(ch), wildcard: true}
					currNode.children = append(currNode.children, node)
				}				
			}
			break
			default:
				if node == nil {
					node = &TrieNode{value: rune(ch)}
					currNode.children = append(currNode.children, node)
				}
				break
		}
		currNode = node
	}
	currNode.isEndofPath = true
	currNode.handler = handler
	currNode.urlpath = path
	currNode.method = method

}

func (t *Trie) Search(path string) (matched_node *TrieNode) {
	currNode := t.root
	for _, wch := range path {
		if wch != '/' && currNode.value == '*' {
			continue
		}
		if currNode = currNode.Find(wch); currNode == nil {
			return nil
		}
	}
	matched_node = currNode
	if !matched_node.wildcard && !matched_node.isEndofPath {
		fmt.Println(string(matched_node.value))
		return nil
	}
	return 
}


func NewTrie() *Trie {
	return &Trie{root: &TrieNode{}}
} 
