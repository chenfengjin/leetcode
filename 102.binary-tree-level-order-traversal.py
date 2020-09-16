#
# @lc app=leetcode id=102 lang=python3
#
# [102] Binary Tree Level Order Traversal
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
# from typing import *
from queue import Queue
class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return []
        # if not root.left and not root.left
        res = []
        current_level_values = []
        current_queue = []
        next_queue = []
        current_queue.append(root)

        while current_queue:
            node = current_queue.pop(0)
            if node.left:
                next_queue.append(node.left)
            if node.right:
                next_queue.append(node.right)
            current_level_values.append(node.val)
            if not current_queue:
                res.append(current_level_values)
                current_level_values = []
                if next_queue:
                    current_queue = next_queue
                    next_queue = []
        
        return res

# @lc code=end

