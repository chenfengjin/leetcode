#
# @lc app=leetcode id=111 lang=python3
#
# [111] Minimum Depth of Binary Tree
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
from typing import *
class Solution: # pay attention to the situation of [1,2]
    def __init__(self):
        self.min_depth = 100000

    def minDepth(self, root: TreeNode) -> int:
        if not root:
            return 0
        
        self.helper(root,1)
        return self.min_depth 

    def helper(self,root,depth):
        if not root.left and not root.right:
            if depth < self.min_depth:
                self.min_depth = depth
            return
        if root.left:
            self.helper(root.left,depth +1 )
        if root.right:
            self.helper(root.right,depth +1)
        
# @lc code=end

