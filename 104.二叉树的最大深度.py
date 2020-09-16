#
# @lc app=leetcode.cn id=104 lang=python3
#
# [104] 二叉树的最大深度
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
        self.max_depth = 0
        # super().__init__()
    def maxDepth(self, root: TreeNode) -> int:
        if not root:
            return 0
        depth = 1
        self.helper(root,depth)
        return self.max_depth
    def helper(self,root,depth):
        if depth > self.max_depth:
            self.max_depth = depth
        if root.left:
            self.helper(root.left,depth+1)
        if root.right:
            self.helper(root.right,depth+1)
# @lc code=end

