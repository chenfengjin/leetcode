/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */
package main

// some keys
// 1. how to identify end of an word
// 2. what data structure to specify the word
// 3. deal with search and startWith
// @lc code=start
type Trie struct {
	nodes map[string]*Trie
}

func Constructor() Trie {
	return Trie{
		nodes: map[string]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		if next, ok := cur.nodes[string(word[i])]; ok {
			cur = next
		} else {
			node := &Trie{
				//  to avoid nil map
				nodes: map[string]*Trie{},
			}
			cur.nodes[string(word[i])] = node
			cur = node
		}
	}
	cur.nodes["/"] = &Trie{}
}

func (this *Trie) Search(word string) bool {
	cur := this
	var ok bool
	for i := 0; i < len(word); i++ {
		if cur, ok = cur.nodes[string(word[i])]; !ok {
			return false
		}
	}
	_, ok = cur.nodes["/"]
	return ok
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	var ok bool
	for i := 0; i < len(prefix); i++ {
		if cur, ok = cur.nodes[string(prefix[i])]; !ok {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
