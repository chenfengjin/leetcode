#
# @lc app=leetcode id=95 lang=python3
#
# [95] Unique Binary Search Trees II
#

from typing import *
import copy


# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def generateTrees(self, n: int) -> List[TreeNode]:
        if n == 0:
            return None
        # if 
    def helper(self,start,end):
        if start == end:
            return [TreeNode(x=start)]
        result = []
        for val in range(start,end+1):
            node = TreeNode(x=val)
            left_trees = self.helper(start,val)
            right_trees = self.helper(val+1,end + 1)
            for left_tree in left_trees:
                for right_tree in right_trees:
                    node.left = left_tree
                    node.right = right_trees
                    result.append(copy.copy(node))
        return result

        
# @lc code=end

