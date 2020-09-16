#
# @lc app=leetcode id=654 lang=python3
#
# [654] Maximum Binary Tree
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def constructMaximumBinaryTree(self, nums: List[int]) -> TreeNode:
        if not nums:
            return None

        index = nums.index(max(nums))
        root = TreeNode(nums[index])
        root.left = self.constructMaximumBinaryTree(nums[:index])
        try:
            root.right = self.constructMaximumBinaryTree(nums[index+1:]) 
        except:
            root.right = None

        return root

        
# @lc code=end

