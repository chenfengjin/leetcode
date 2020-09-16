#
# @lc app=leetcode id=189 lang=python3
#
# [189] Rotate Array
#
from typing import *
# @lc code=start
class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        if not nums or len(nums) ==1 :
            return 
        # tmp = 0
        # for _ in range(k):
        #     tmp = nums[-1]
        #     for index in range(len(nums)-1,0,-1):
        #         nums[index] = nums[index-1]
        #     nums[0] = tmp
        left = nums[-1*k]
        right = nums[:k+1]
        length = len(nums)
        nums[:length-k] = right
        nums[length-k+1:] = left


        
# @lc code=end

print(Solution().rotate([1,2,3,4,5,6,7],3))