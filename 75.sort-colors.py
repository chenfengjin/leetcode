#
# @lc app=leetcode id=75 lang=python3
#
# [75] Sort Colors
#
import copy
from typing import *
from collections import deque
# @lc code=start
class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        left  = 0
        right = len(nums) -1
        index  = 0
        while index <= right:
            if nums[index] == 0:
                nums[left],nums[index] = nums[index],nums[left]
                index  += 1
                left += 1
            elif nums[index] == 1:
                index += 1
            else:
                nums[right],nums[index] = nums[index],nums[right]
                right -= 1

# @lc code=end
# 1 2 3 4 5
# 0 1 -3 -2 -2 
if __name__ == "__main__":
    print(Solution().sortColors([1,2,0]))