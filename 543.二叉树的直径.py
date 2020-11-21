#
# @lc app=leetcode.cn id=543 lang=python3
#
# [543] 二叉树的直径
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def __init__(self):
        self.res = 0
    def diameterOfBinaryTree(self, root: TreeNode) -> int:
        if not root:
            return -1
        l = self.diameterOfBinaryTree(root.left)
        r = self.diameterOfBinaryTree(root.right)
        if xx:
            self.res = xx
        return xx
# @lc code=end

