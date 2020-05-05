#
# @lc app=leetcode id=108 lang=python3
#
# [108] Convert Sorted Array to Binary Search Tree
#

# @lc code=start
# Definition for a binary tree node.

from typing import *
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> TreeNode:
        if not nums:
            return None
        return self.helper(nums)

    def helper(self,nums):
        if len(nums) == 1:
            return TreeNode(nums[0])
        middle = (len(nums))//2
        root =  TreeNode(nums[middle])
        root.left = self.helper(nums[0:middle])
        root.right = self.helper(nums[middle+1:]) if middle + 1 < len(nums) else None
        return root       
# @lc code=end

