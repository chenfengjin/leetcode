#
# @lc app=leetcode id=95 lang=python3
#
# [95] Unique Binary Search Trees II
#
from typing import *
import copy
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
# @lc code=start
# too slow
class Solution:
    def __init__(self):
        self.memory = {}
    def generateTrees(self, n: int) -> List[TreeNode]:
        self.memory = [[0] * (n +1)] * (n +1)
        return  self.generateSubTree(list(range(1,n+1)))    
    def generateSubTree(self,nums):
        if not nums:  
            return [] # key point None is an object of NoneType
        if len(nums) == 1:
            return [TreeNode(nums[0])]

        # if self.memory[nums[0]][nums[-1]]:
            # return copy.deepcopy(self.memory[nums[0]][nums[-1]])
        result = []
        for i in range(len(nums)):
            left_subtrees = self.generateSubTree(nums[:i])
            right_subtrees = self.generateSubTree(nums[i+1:]) if i+1 < len(nums) else []  # key point 2
            root = TreeNode(nums[i])
            if left_subtrees:
                if right_subtrees:
                    for left_tree in left_subtrees:
                        for right_tree in right_subtrees:
                            root.left = left_tree
                            root.right = right_tree
                            result.append(copy.deepcopy(root))                   
                else:
                    for left_tree in left_subtrees:
                        root.left = left_tree
                        result.append(copy.deepcopy(root))                   
                        
            else:
                if right_subtrees:
                    for right_tree in right_subtrees:
                        root.right = right_tree
                        result.append(copy.deepcopy(root))                   
                else:
                    result.append(root)
        # self.memory[nums[0]][nums[-1]] = copy.deepcopy(result)
        return result
# @lc code=end

if __name__ == "__main__":
    print((Solution().generateTrees(4)))