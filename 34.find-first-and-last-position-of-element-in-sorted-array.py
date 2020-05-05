#
# @lc app=leetcode id=34 lang=python3
#
# [34] Find First and Last Position of Element in Sorted Array
#

# @lc code=start
class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        left = 0 
        right = len(nums)
        while left <= right:
            middle = (left+ right) // 2
            if nums[middle] ==  target:
                return middle
            if nums[middle] > target:
                right = middle - 1
            else:
                left = middle + 1
        
# @lc code=end

