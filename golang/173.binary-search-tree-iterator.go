/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nodes []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		nodes: []*TreeNode{root},
	}

}

func (this *BSTIterator) Next() int {
	node := this.nodes[len(this.nodes)-1]
	this.nodes = this.nodes[:len(this.nodes)-1]
	left := node.Left
	right := node.Right

	node.Left = nil
	node.Right = nil

	if right != nil {
		this.nodes = append(this.nodes, right)
	}
	if left == nil {
		return node.Val
	} else {
		this.nodes = append(this.nodes, node)
		this.nodes = append(this.nodes, left)
		return this.Next()
	}
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) != 0

}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
