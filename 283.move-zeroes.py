#
# @lc app=leetcode id=283 lang=python3
#
# [283] Move Zeroes
#

from typing import *
# @lc code=start
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        # 操作数是最少吗?
        left = right = 0
        while right < len(nums):
            if nums[left] == 0:
                nums[left] = nums[right]
                right += 1
            else:
                left += 1
        n
                            
# @lc code=end
