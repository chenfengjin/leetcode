/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */
package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 1. 层顺遍历，然后判断每层是否是回文，实现见
// 1.1 层序遍历见 102 103
// 1.2 回文判断见 125
// 2. 判断左子树和右子树是否镜像
// 边界条件： 空树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)

}
func isMirror(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)

}

// @lc code=end
