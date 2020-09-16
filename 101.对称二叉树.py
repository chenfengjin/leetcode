#
# @lc app=leetcode.cn id=101 lang=python3
#
# [101] 对称二叉树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool: #递归解法
        if not root:
            return True
        if not root.left and not root.right:
            return True
        if root.left and root.right:
            return root.left.val == root.right.val and self.isSymmetric(root.left) and self.isSymmetric(root.right)
        return False        
# @lc code=end

