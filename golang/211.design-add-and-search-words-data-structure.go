/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */
//  differ from 208 implement tree, we use recursion to implement search, which make code easy to understand
package main

// TODO
//  have enough motivation to make it run faster
// @lc code=start
type WordDictionary struct {
	nodes map[string]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{
		nodes: map[string]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		if next, ok := cur.nodes[string(word[i])]; ok {
			cur = next
		} else {
			node := &WordDictionary{
				//  to avoid nil map
				nodes: map[string]*WordDictionary{},
			}
			cur.nodes[string(word[i])] = node
			cur = node
		}
	}
	cur.nodes["/"] = &WordDictionary{}
}

func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		_, ok := this.nodes["/"]
		return ok
	}
	if word[0] == '.' {
		for _, nodes := range this.nodes {
			if nodes.Search(word[1:]) {
				return true
			}
		}
		return false
	}
	node, ok := this.nodes[string(word[0])]
	if !ok {
		return false
	}
	return node.Search(word[1:])
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end
