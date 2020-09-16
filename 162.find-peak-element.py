#
# @lc app=leetcode id=162 lang=python3
#
# [162] Find Peak Element
#

# @lc code=start
class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 0
        if nums[0] > nums[1]:
            return 0
        if nums[-1] > nums[-2]:
            return len(nums) - 1
        for i in range(1,len(nums)-1): # 这里可以剪枝
            if nums[i+1] < nums[i] and nums[i] > nums[i-1]:
                return i
        
# @lc code=end

